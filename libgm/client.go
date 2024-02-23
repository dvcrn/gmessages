package libgm

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"go.mau.fi/mautrix-gmessages/libgm/crypto"
	"go.mau.fi/mautrix-gmessages/libgm/events"
	"go.mau.fi/mautrix-gmessages/libgm/gmproto"
	"go.mau.fi/mautrix-gmessages/libgm/util"
)

type AuthData struct {
	// Keys used to encrypt communication with the phone
	RequestCrypto *crypto.AESCTRHelper `json:"request_crypto,omitempty"`
	// Key used to sign requests to refresh the tachyon auth token from the server
	RefreshKey *crypto.JWK `json:"refresh_key,omitempty"`
	// Identity of the paired phone and browser
	Browser *gmproto.Device `json:"browser,omitempty"`
	Mobile  *gmproto.Device `json:"mobile,omitempty"`
	// Key used to authenticate with the server
	TachyonAuthToken []byte    `json:"tachyon_token,omitempty"`
	TachyonExpiry    time.Time `json:"tachyon_expiry,omitempty"`
	TachyonTTL       int64     `json:"tachyon_ttl,omitempty"`
	// Unknown encryption key, not used for anything
	WebEncryptionKey []byte `json:"web_encryption_key,omitempty"`

	SessionID uuid.UUID         `json:"session_id,omitempty"`
	Cookies   map[string]string `json:"cookies,omitempty"`
}

func (ad *AuthData) AuthNetwork() string {
	if ad.Cookies != nil {
		return util.GoogleNetwork
	}
	return ""
}

const RefreshTachyonBuffer = 1 * time.Hour

type Proxy func(*http.Request) (*url.URL, error)
type EventHandler func(evt any)

type updateDedupItem struct {
	id   string
	hash [32]byte
}

type Client struct {
	Logger         zerolog.Logger
	evHandler      EventHandler
	sessionHandler *SessionHandler

	longPollingConn io.Closer
	listenID        int
	skipCount       int
	disconnecting   bool

	pingShortCircuit chan struct{}

	recentUpdates    [8]updateDedupItem
	recentUpdatesPtr int

	conversationsFetchedOnce bool

	AuthData *AuthData
	cfg      *gmproto.Config

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
		responseWaiters: make(map[string]chan<- *IncomingRPCMessage),
	}
	cli := &Client{
		AuthData:       authData,
		Logger:         logger,
		sessionHandler: sessionHandler,
		http:           &http.Client{},

		pingShortCircuit: make(chan struct{}),
	}
	sessionHandler.client = cli
	var err error
	cli.cfg, err = cli.FetchConfig()
	if err != nil {
		cli.Logger.Err(err).Msg("Failed to fetch web config")
	} else if deviceID := cli.cfg.GetDeviceInfo().GetDeviceID(); deviceID != "" {
		authData.SessionID, err = uuid.Parse(deviceID)
		if err != nil {
			cli.Logger.Err(err).Str("device_id", deviceID).Msg("Failed to parse device ID")
		}
	}
	return cli
}

func (c *Client) CurrentSessionID() string {
	return c.sessionHandler.sessionID
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
	if c.AuthData.TachyonAuthToken == nil {
		return fmt.Errorf("no auth token")
	} else if c.AuthData.Browser == nil {
		return fmt.Errorf("not logged in")
	}

	err := c.refreshAuthToken()
	if err != nil {
		return fmt.Errorf("failed to refresh auth token: %w", err)
	}

	//webEncryptionKeyResponse, err := c.GetWebEncryptionKey()
	//if err != nil {
	//	return fmt.Errorf("failed to get web encryption key: %w", err)
	//}
	//c.updateWebEncryptionKey(webEncryptionKeyResponse.GetKey())
	go c.doLongPoll(true)
	c.sessionHandler.startAckInterval()
	go c.postConnect()
	return nil
}

func (c *Client) postConnect() {
	err := c.SetActiveSession()
	if err != nil {
		c.Logger.Err(err).Msg("Failed to set active session")
		c.triggerEvent(&events.PingFailed{
			Error: fmt.Errorf("failed to set active session: %w", err),
		})
		return
	}
	if c.AuthData.Mobile.Network != util.QRNetwork {
		// Don't check bugle default unless using bugle
		return
	}

	doneChan := make(chan struct{})
	go func() {
		select {
		case <-time.After(5 * time.Second):
			c.Logger.Warn().Msg("Checking bugle default on connect is taking long")
		case <-doneChan:
		}
	}()
	bugleRes, err := c.IsBugleDefault()
	close(doneChan)
	if err != nil {
		c.Logger.Err(err).Msg("Failed to check bugle default")
		return
	}
	c.Logger.Debug().Bool("bugle_default", bugleRes.Success).Msg("Got is bugle default response on connect")
}

func (c *Client) Disconnect() {
	c.closeLongPolling()
	c.http.CloseIdleConnections()
}

func (c *Client) IsConnected() bool {
	// TODO add better check (longPollingConn is set to nil while the polling reconnects)
	return c.longPollingConn != nil
}

func (c *Client) IsLoggedIn() bool {
	return c.AuthData != nil && c.AuthData.Browser != nil
}

func (c *Client) Reconnect() error {
	c.closeLongPolling()
	err := c.Connect()
	if err != nil {
		c.Logger.Err(err).Msg("Failed to reconnect")
		return err
	}
	c.Logger.Debug().Msg("Successfully reconnected to server")
	return nil
}

func (c *Client) triggerEvent(evt interface{}) {
	if c.evHandler != nil {
		c.evHandler(evt)
	}
}

func (c *Client) FetchConfig() (*gmproto.Config, error) {
	req, err := http.NewRequest(http.MethodGet, util.ConfigURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %w", err)
	}
	util.BuildRelayHeaders(req, "", "*/*")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Del("x-user-agent")
	req.Header.Del("origin")
	c.AddCookieHeaders(req)

	resp, err := c.http.Do(req)
	if resp != nil {
		c.HandleCookieUpdates(resp)
	}
	config, err := typedHTTPResponse[*gmproto.Config](resp, err)
	if err != nil {
		return nil, err
	}

	version, parseErr := config.ParsedClientVersion()
	if parseErr != nil {
		return nil, fmt.Errorf("failed to parse client version: %w", err)
	}

	currVersion := util.ConfigMessage
	if version.Year != currVersion.Year || version.Month != currVersion.Month || version.Day != currVersion.Day {
		toLog := c.diffVersionFormat(currVersion, version)
		c.Logger.Trace().Any("version", toLog).Msg("Messages for web version is not latest")
	} else {
		c.Logger.Debug().Any("version", currVersion).Msg("Using latest messages for web version")
	}

	return config, nil
}

func (c *Client) diffVersionFormat(curr *gmproto.ConfigVersion, latest *gmproto.ConfigVersion) string {
	return fmt.Sprintf("%d.%d.%d -> %d.%d.%d", curr.Year, curr.Month, curr.Day, latest.Year, latest.Month, latest.Day)
}

func (c *Client) updateWebEncryptionKey(key []byte) {
	c.Logger.Debug().Msg("Updated WebEncryptionKey")
	c.AuthData.WebEncryptionKey = key
}

func (c *Client) updateTachyonAuthToken(data *gmproto.TokenData) {
	c.AuthData.TachyonAuthToken = data.GetTachyonAuthToken()
	validForDuration := time.Duration(data.GetTTL()) * time.Microsecond
	if validForDuration == 0 {
		validForDuration = 24 * time.Hour
	}
	c.AuthData.TachyonExpiry = time.Now().UTC().Add(validForDuration)
	c.AuthData.TachyonTTL = validForDuration.Microseconds()
	c.Logger.Debug().
		Time("tachyon_expiry", c.AuthData.TachyonExpiry).
		Int64("valid_for", data.GetTTL()).
		Msg("Updated tachyon token")
}

func (c *Client) refreshAuthToken() error {
	if c.AuthData.Browser == nil || time.Until(c.AuthData.TachyonExpiry) > RefreshTachyonBuffer {
		return nil
	}
	c.Logger.Debug().
		Time("tachyon_expiry", c.AuthData.TachyonExpiry).
		Msg("Refreshing auth token")
	jwk := c.AuthData.RefreshKey
	requestID := uuid.NewString()
	timestamp := time.Now().UnixMilli() * 1000

	signBytes := sha256.Sum256([]byte(fmt.Sprintf("%s:%d", requestID, timestamp)))
	sig, err := ecdsa.SignASN1(rand.Reader, jwk.GetPrivateKey(), signBytes[:])
	if err != nil {
		return err
	}

	payload := &gmproto.RegisterRefreshRequest{
		MessageAuth: &gmproto.AuthMessage{
			RequestID:        requestID,
			TachyonAuthToken: c.AuthData.TachyonAuthToken,
			Network:          c.AuthData.AuthNetwork(),
			ConfigVersion:    util.ConfigMessage,
		},
		CurrBrowserDevice: c.AuthData.Browser,
		UnixTimestamp:     timestamp,
		Signature:         sig,
		EmptyRefreshArr:   &gmproto.RegisterRefreshRequest_NestedEmptyArr{EmptyArr: &gmproto.EmptyArr{}},
		MessageType:       2, // hmm
	}

	resp, err := typedHTTPResponse[*gmproto.RegisterRefreshResponse](
		c.makeProtobufHTTPRequest(util.RegisterRefreshURL, payload, ContentTypePBLite),
	)
	if err != nil {
		return err
	}

	if resp.GetTokenData().GetTachyonAuthToken() == nil {
		return fmt.Errorf("no tachyon auth token in refresh response")
	}

	c.updateTachyonAuthToken(resp.GetTokenData())
	c.triggerEvent(&events.AuthTokenRefreshed{})
	return nil
}
