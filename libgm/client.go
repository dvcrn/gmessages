package libgm

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"

	"go.mau.fi/mautrix-gmessages/libgm/binary"
	"go.mau.fi/mautrix-gmessages/libgm/crypto"
	"go.mau.fi/mautrix-gmessages/libgm/events"
	"go.mau.fi/mautrix-gmessages/libgm/payload"
	"go.mau.fi/mautrix-gmessages/libgm/pblite"
	"go.mau.fi/mautrix-gmessages/libgm/util"
)

type AuthData struct {
	// Keys used to encrypt communication with the phone
	RequestCrypto *crypto.AESCTRHelper `json:"request_crypto,omitempty"`
	// Key used to sign requests to refresh the tachyon auth token from the server
	RefreshKey *crypto.JWK `json:"refresh_key,omitempty"`
	// Identity of the paired phone and browser
	Browser *binary.Device `json:"browser,omitempty"`
	Mobile  *binary.Device `json:"mobile,omitempty"`
	// Key used to authenticate with the server
	TachyonAuthToken []byte    `json:"tachyon_token,omitempty"`
	TachyonExpiry    time.Time `json:"tachyon_expiry,omitempty"`
	TachyonTTL       int64     `json:"tachyon_ttl,omitempty"`
	// Unknown encryption key, not used for anything
	WebEncryptionKey []byte `json:"web_encryption_key,omitempty"`
}

const RefreshTachyonBuffer = 1 * time.Hour

type Proxy func(*http.Request) (*url.URL, error)
type EventHandler func(evt any)

type Client struct {
	Logger         zerolog.Logger
	rpc            *RPC
	pairer         *Pairer
	evHandler      EventHandler
	sessionHandler *SessionHandler

	authData *AuthData

	proxy Proxy
	http  *http.Client
}

func NewAuthData() *AuthData {
	return &AuthData{
		RequestCrypto: crypto.NewAESCTRHelper(),
		RefreshKey:    crypto.GenerateECDSAKey(),
	}
}

func NewClient(authData *AuthData, logger zerolog.Logger) *Client {
	sessionHandler := &SessionHandler{
		responseWaiters: make(map[string]chan<- *pblite.Response),
		responseTimeout: time.Duration(5000) * time.Millisecond,
	}
	cli := &Client{
		authData:       authData,
		Logger:         logger,
		sessionHandler: sessionHandler,
		http:           &http.Client{},
	}
	sessionHandler.client = cli
	rpc := &RPC{client: cli, http: &http.Client{Transport: &http.Transport{Proxy: cli.proxy}}}
	cli.rpc = rpc
	cli.FetchConfigVersion()
	return cli
}

func (c *Client) SetEventHandler(eventHandler EventHandler) {
	c.evHandler = eventHandler
}

func (c *Client) SetProxy(proxy string) error {
	proxyParsed, err := url.Parse(proxy)
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("Failed to set proxy")
	}
	proxyUrl := http.ProxyURL(proxyParsed)
	c.http.Transport = &http.Transport{
		Proxy: proxyUrl,
	}
	c.proxy = proxyUrl
	c.Logger.Debug().Any("proxy", proxyParsed.Host).Msg("SetProxy")
	return nil
}
func (c *Client) Connect() error {
	if c.authData.TachyonAuthToken != nil {
		refreshErr := c.refreshAuthToken()
		if refreshErr != nil {
			panic(refreshErr)
		}

		webEncryptionKeyResponse, webEncryptionKeyErr := c.GetWebEncryptionKey()
		if webEncryptionKeyErr != nil {
			c.Logger.Err(webEncryptionKeyErr).Any("response", webEncryptionKeyResponse).Msg("GetWebEncryptionKey request failed")
			return webEncryptionKeyErr
		}
		c.updateWebEncryptionKey(webEncryptionKeyResponse.GetKey())
		go c.rpc.ListenReceiveMessages()
		c.sessionHandler.startAckInterval()

		bugleRes, bugleErr := c.IsBugleDefault()
		if bugleErr != nil {
			panic(bugleErr)
		}
		c.Logger.Info().Any("isBugle", bugleRes.Success).Msg("IsBugleDefault")
		sessionErr := c.SetActiveSession()
		if sessionErr != nil {
			panic(sessionErr)
		}
		return nil
	} else {
		pairer, err := c.NewPairer(c.authData.RefreshKey, 20)
		if err != nil {
			panic(err)
		}
		c.pairer = pairer
		registered, err2 := c.pairer.RegisterPhoneRelay()
		if err2 != nil {
			return err2
		}
		c.authData.TachyonAuthToken = registered.AuthKeyData.TachyonAuthToken
		go c.rpc.ListenReceiveMessages()
		return nil
	}
}

func (c *Client) Disconnect() {
	c.rpc.CloseConnection()
	c.http.CloseIdleConnections()
}

func (c *Client) IsConnected() bool {
	return c.rpc != nil
}

func (c *Client) IsLoggedIn() bool {
	return c.authData != nil && c.authData.Browser != nil
}

func (c *Client) Reconnect() error {
	c.rpc.CloseConnection()
	for c.rpc.conn != nil {
		time.Sleep(time.Millisecond * 100)
	}
	err := c.Connect()
	if err != nil {
		c.Logger.Err(err).Any("tachyonAuthToken", c.authData.TachyonAuthToken).Msg("Failed to reconnect")
		return err
	}
	c.Logger.Debug().Any("tachyonAuthToken", c.authData.TachyonAuthToken).Msg("Successfully reconnected to server")
	return nil
}

func (c *Client) triggerEvent(evt interface{}) {
	if c.evHandler != nil {
		c.evHandler(evt)
	}
}

func (c *Client) DownloadMedia(mediaID string, key []byte) ([]byte, error) {
	downloadMetadata := &binary.UploadImagePayload{
		MetaData: &binary.ImageMetaData{
			ImageID:   mediaID,
			Encrypted: true,
		},
		AuthData: &binary.AuthMessage{
			RequestID:        uuid.NewString(),
			TachyonAuthToken: c.authData.TachyonAuthToken,
			ConfigVersion:    payload.ConfigMessage,
		},
	}
	downloadMetadataBytes, err2 := proto.Marshal(downloadMetadata)
	if err2 != nil {
		return nil, err2
	}
	downloadMetadataEncoded := base64.StdEncoding.EncodeToString(downloadMetadataBytes)
	req, err := http.NewRequest("GET", util.UploadMediaURL, nil)
	if err != nil {
		return nil, err
	}
	util.BuildUploadHeaders(req, downloadMetadataEncoded)
	res, reqErr := c.http.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}
	c.Logger.Info().Any("url", util.UploadMediaURL).Any("headers", res.Request.Header).Msg("Decrypt Image Headers")
	defer res.Body.Close()
	encryptedBuffImg, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		return nil, err3
	}
	c.Logger.Debug().Any("key", key).Any("encryptedLength", len(encryptedBuffImg)).Msg("Attempting to decrypt image")
	cryptor, err := crypto.NewImageCryptor(key)
	if err != nil {
		return nil, err
	}
	decryptedImageBytes, decryptionErr := cryptor.DecryptData(encryptedBuffImg)
	if decryptionErr != nil {
		return nil, decryptionErr
	}
	return decryptedImageBytes, nil
}

func (c *Client) FetchConfigVersion() {
	req, bErr := http.NewRequest("GET", util.ConfigUrl, nil)
	if bErr != nil {
		panic(bErr)
	}

	configRes, requestErr := c.http.Do(req)
	if requestErr != nil {
		panic(requestErr)
	}

	responseBody, readErr := io.ReadAll(configRes.Body)
	if readErr != nil {
		panic(readErr)
	}

	version, parseErr := util.ParseConfigVersion(responseBody)
	if parseErr != nil {
		panic(parseErr)
	}

	currVersion := payload.ConfigMessage
	if version.Year != currVersion.Year || version.Month != currVersion.Month || version.Day != currVersion.Day {
		toLog := c.diffVersionFormat(currVersion, version)
		c.Logger.Info().Any("version", toLog).Msg("There's a new version available!")
	} else {
		c.Logger.Info().Any("version", currVersion).Msg("You are running on the latest version.")
	}
}

func (c *Client) diffVersionFormat(curr *binary.ConfigVersion, latest *binary.ConfigVersion) string {
	return fmt.Sprintf("%d.%d.%d -> %d.%d.%d", curr.Year, curr.Month, curr.Day, latest.Year, latest.Month, latest.Day)
}

func (c *Client) updateWebEncryptionKey(key []byte) {
	c.Logger.Debug().Any("key", key).Msg("Updated WebEncryptionKey")
	c.authData.WebEncryptionKey = key
}

func (c *Client) updateTachyonAuthToken(t []byte, validFor int64) {
	c.authData.TachyonAuthToken = t
	validForDuration := time.Duration(validFor) * time.Microsecond
	if validForDuration == 0 {
		validForDuration = 24 * time.Hour
	}
	c.authData.TachyonExpiry = time.Now().UTC().Add(time.Microsecond * time.Duration(validFor))
	c.authData.TachyonTTL = validForDuration.Microseconds()
	c.Logger.Debug().Time("tachyon_expiry", c.authData.TachyonExpiry).Int64("valid_for", validFor).Msg("Updated tachyon token")
}

func (c *Client) updateDevicePair(mobile, browser *binary.Device) {
	c.authData.Mobile = mobile
	c.authData.Browser = browser
	c.Logger.Debug().Any("mobile", mobile).Any("browser", browser).Msg("Updated device pair")
}

func (c *Client) SaveAuthSession(path string) error {
	toSaveJson, jsonErr := json.Marshal(c.authData)
	if jsonErr != nil {
		return jsonErr
	}
	writeErr := os.WriteFile(path, toSaveJson, os.ModePerm)
	return writeErr
}

func LoadAuthSession(path string) (*AuthData, error) {
	jsonData, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}

	sessionData := &AuthData{}
	marshalErr := json.Unmarshal(jsonData, sessionData)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return sessionData, nil
}

func (c *Client) refreshAuthToken() error {
	if c.authData.Browser == nil || time.Until(c.authData.TachyonExpiry) > RefreshTachyonBuffer {
		return nil
	}
	c.Logger.Debug().Time("tachyon_expiry", c.authData.TachyonExpiry).Msg("Refreshing auth token")
	jwk := c.authData.RefreshKey
	requestID := uuid.NewString()
	timestamp := time.Now().UnixMilli() * 1000

	signBytes := sha256.Sum256([]byte(fmt.Sprintf("%s:%d", requestID, timestamp)))
	sig, err := ecdsa.SignASN1(rand.Reader, jwk.GetPrivateKey(), signBytes[:])
	if err != nil {
		return err
	}

	payload, err := pblite.Marshal(&binary.RegisterRefreshPayload{
		MessageAuth: &binary.AuthMessage{
			RequestID:        requestID,
			TachyonAuthToken: c.authData.TachyonAuthToken,
			ConfigVersion:    payload.ConfigMessage,
		},
		CurrBrowserDevice: c.authData.Browser,
		UnixTimestamp:     timestamp,
		Signature:         sig,
		EmptyRefreshArr:   &binary.EmptyRefreshArr{EmptyArr: &binary.EmptyArr{}},
		MessageType:       2, // hmm
	})
	if err != nil {
		return err
	}

	refreshResponse, requestErr := c.rpc.sendMessageRequest(util.RegisterRefreshURL, payload)
	if requestErr != nil {
		return requestErr
	}

	if refreshResponse.StatusCode == 401 {
		return fmt.Errorf("failed to refresh auth token: unauthorized (try reauthenticating through qr code)")
	}

	if refreshResponse.StatusCode == 400 {
		return fmt.Errorf("failed to refresh auth token: signature failed")
	}
	responseBody, readErr := io.ReadAll(refreshResponse.Body)
	if readErr != nil {
		return readErr
	}

	resp := &binary.RegisterRefreshResponse{}
	deserializeErr := pblite.Unmarshal(responseBody, resp)
	if deserializeErr != nil {
		return deserializeErr
	}

	token := resp.GetTokenData().GetTachyonAuthToken()
	if token == nil {
		return fmt.Errorf("failed to refresh auth token: something happened")
	}

	validFor, _ := strconv.ParseInt(resp.GetTokenData().GetValidFor(), 10, 64)

	c.updateTachyonAuthToken(token, validFor)
	c.triggerEvent(events.NewAuthTokenRefreshed(token))
	return nil
}
