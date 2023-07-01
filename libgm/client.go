package libgm

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog"

	"go.mau.fi/mautrix-gmessages/libgm/binary"
	"go.mau.fi/mautrix-gmessages/libgm/crypto"
	"go.mau.fi/mautrix-gmessages/libgm/payload"
	"go.mau.fi/mautrix-gmessages/libgm/util"
)

type DevicePair struct {
	Mobile  *binary.Device
	Browser *binary.Device
}
type Proxy func(*http.Request) (*url.URL, error)
type EventHandler func(evt interface{})
type Client struct {
	Logger         zerolog.Logger
	Conversations  *Conversations
	Session        *Session
	rpc            *RPC
	devicePair     *DevicePair
	pairer         *Pairer
	cryptor        *crypto.Cryptor
	imageCryptor   *crypto.ImageCryptor
	evHandler      EventHandler
	sessionHandler *SessionHandler
	instructions   *Instructions

	rpcKey []byte
	ttl    int64

	proxy Proxy
	http  *http.Client
}

func NewClient(devicePair *DevicePair, cryptor *crypto.Cryptor, logger zerolog.Logger, proxy *string) *Client {
	sessionHandler := &SessionHandler{
		requests:        make(map[string]map[int64]*ResponseChan),
		responseTimeout: time.Duration(5000) * time.Millisecond,
	}
	if cryptor == nil {
		cryptor = crypto.NewCryptor(nil, nil)
	}
	cli := &Client{
		Logger:         logger,
		devicePair:     devicePair,
		sessionHandler: sessionHandler,
		cryptor:        cryptor,
		imageCryptor:   &crypto.ImageCryptor{},
		http:           &http.Client{},
	}
	sessionHandler.client = cli
	cli.instructions = NewInstructions(cli.cryptor)
	if proxy != nil {
		cli.SetProxy(*proxy)
	}
	rpc := &RPC{client: cli, http: &http.Client{Transport: &http.Transport{Proxy: cli.proxy}}}
	cli.rpc = rpc
	cli.Logger.Debug().Any("data", cryptor).Msg("Cryptor")
	cli.setApiMethods()
	return cli
}

func (c *Client) SetEventHandler(eventHandler EventHandler) {
	if eventHandler == nil {
		eventHandler = func(_ interface{}) {}
	}
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

func (c *Client) Connect(rpcKey []byte) error {
	rpcPayload, receiveMesageSessionID, err := payload.ReceiveMessages(rpcKey)
	if err != nil {
		panic(err)
		return err
	}
	c.rpc.rpcSessionID = receiveMesageSessionID
	c.rpcKey = rpcKey
	go c.rpc.ListenReceiveMessages(rpcPayload)
	c.Logger.Debug().Any("rpcKey", rpcKey).Msg("Successfully connected to server")
	if c.devicePair != nil {
		sendInitialDataErr := c.rpc.sendInitialData()
		if sendInitialDataErr != nil {
			panic(sendInitialDataErr)
		}
	}
	return nil
}

func (c *Client) Reconnect(rpcKey []byte) error {
	c.rpc.CloseConnection()
	for c.rpc.conn != nil {
		time.Sleep(time.Millisecond * 100)
	}
	err := c.Connect(rpcKey)
	if err != nil {
		c.Logger.Err(err).Any("rpcKey", rpcKey).Msg("Failed to reconnect")
		return err
	}
	c.Logger.Debug().Any("rpcKey", rpcKey).Msg("Successfully reconnected to server")
	sendInitialDataErr := c.rpc.sendInitialData()
	if sendInitialDataErr != nil {
		panic(sendInitialDataErr)
	}
	return nil
}

func (c *Client) Disconnect() {
	c.rpc.CloseConnection()
	c.http.CloseIdleConnections()
}

func (c *Client) IsConnected() bool {
	return c.rpc != nil
}

func (c *Client) IsLoggedIn() bool {
	return c.devicePair != nil
}

func (c *Client) triggerEvent(evt interface{}) {
	if c.evHandler != nil {
		c.evHandler(evt)
	}
}

func (c *Client) setApiMethods() {
	c.Conversations = &Conversations{
		client: c,
		openConversation: openConversation{
			client: c,
		},
		fetchConversationMessages: fetchConversationMessages{
			client: c,
		},
	}
	c.Session = &Session{
		client: c,
		prepareNewSession: prepareNewSession{
			client: c,
		},
		newSession: newSession{
			client: c,
		},
	}
}

func (c *Client) decryptImages(messages *binary.FetchMessagesResponse) error {
	for _, msg := range messages.Messages {
		switch msg.GetType() {
		case *binary.MessageType_IMAGE.Enum():
			for _, details := range msg.GetMessageInfo() {
				switch data := details.GetData().(type) {
				case *binary.MessageInfo_ImageContent:
					decryptedImageData, err := c.decryptImageData(data.ImageContent.ImageID, data.ImageContent.DecryptionKey)
					if err != nil {
						panic(err)
						return err
					}
					data.ImageContent.ImageData = decryptedImageData
				}
			}
		}
	}
	return nil
}

func (c *Client) decryptImageData(imageId string, key []byte) ([]byte, error) {
	reqId := util.RandomUUIDv4()
	download_metadata := &binary.UploadImagePayload{
		MetaData: &binary.ImageMetaData{
			ImageID:   imageId,
			Encrypted: true,
		},
		AuthData: &binary.AuthMessage{
			RequestID: reqId,
			RpcKey:    c.rpcKey,
			Date: &binary.Date{
				Year: 2023,
				Seq1: 6,
				Seq2: 8,
				Seq3: 4,
				Seq4: 6,
			},
		},
	}
	download_metadata_bytes, err2 := binary.EncodeProtoMessage(download_metadata)
	if err2 != nil {
		return nil, err2
	}
	download_metadata_b64 := base64.StdEncoding.EncodeToString(download_metadata_bytes)
	req, err := http.NewRequest("GET", util.UPLOAD_MEDIA, nil)
	if err != nil {
		return nil, err
	}
	util.BuildUploadHeaders(req, download_metadata_b64)
	res, reqErr := c.http.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}
	c.Logger.Info().Any("url", util.UPLOAD_MEDIA).Any("headers", res.Request.Header).Msg("Decrypt Image Headers")
	defer res.Body.Close()
	encryptedBuffImg, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		return nil, err3
	}
	c.Logger.Debug().Any("key", key).Any("encryptedLength", len(encryptedBuffImg)).Msg("Attempting to decrypt image")
	c.imageCryptor.UpdateDecryptionKey(key)
	decryptedImageBytes, decryptionErr := c.imageCryptor.DecryptData(encryptedBuffImg)
	if decryptionErr != nil {
		c.Logger.Err(err).Msg("Image decryption failed")
		return nil, decryptionErr
	}
	return decryptedImageBytes, nil
}
