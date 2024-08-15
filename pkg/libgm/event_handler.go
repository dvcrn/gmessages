package libgm

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"

	"go.mau.fi/mautrix-gmessages/pkg/libgm/events"
	"go.mau.fi/mautrix-gmessages/pkg/libgm/gmproto"
)

type IncomingRPCMessage struct {
	*gmproto.IncomingRPCMessage

	IsOld bool

	Pair *gmproto.RPCPairData
	Gaia *gmproto.RPCGaiaData

	Message          *gmproto.RPCMessageData
	DecryptedData    []byte
	DecryptedMessage proto.Message
}

var responseType = map[gmproto.ActionType]proto.Message{
	gmproto.ActionType_IS_BUGLE_DEFAULT:           &gmproto.IsBugleDefaultResponse{},
	gmproto.ActionType_GET_UPDATES:                &gmproto.UpdateEvents{},
	gmproto.ActionType_LIST_CONVERSATIONS:         &gmproto.ListConversationsResponse{},
	gmproto.ActionType_NOTIFY_DITTO_ACTIVITY:      &gmproto.NotifyDittoActivityResponse{},
	gmproto.ActionType_GET_CONVERSATION_TYPE:      &gmproto.GetConversationTypeResponse{},
	gmproto.ActionType_GET_CONVERSATION:           &gmproto.GetConversationResponse{},
	gmproto.ActionType_LIST_MESSAGES:              &gmproto.ListMessagesResponse{},
	gmproto.ActionType_SEND_MESSAGE:               &gmproto.SendMessageResponse{},
	gmproto.ActionType_SEND_REACTION:              &gmproto.SendReactionResponse{},
	gmproto.ActionType_DELETE_MESSAGE:             &gmproto.DeleteMessageResponse{},
	gmproto.ActionType_GET_PARTICIPANTS_THUMBNAIL: &gmproto.GetThumbnailResponse{},
	gmproto.ActionType_GET_CONTACTS_THUMBNAIL:     &gmproto.GetThumbnailResponse{},
	gmproto.ActionType_LIST_CONTACTS:              &gmproto.ListContactsResponse{},
	gmproto.ActionType_LIST_TOP_CONTACTS:          &gmproto.ListTopContactsResponse{},
	gmproto.ActionType_GET_OR_CREATE_CONVERSATION: &gmproto.GetOrCreateConversationResponse{},
	gmproto.ActionType_UPDATE_CONVERSATION:        &gmproto.UpdateConversationResponse{},
	gmproto.ActionType_GET_FULL_SIZE_IMAGE:        &gmproto.GetFullSizeImageResponse{},
}

func (c *Client) decryptInternalMessage(data *gmproto.IncomingRPCMessage) (*IncomingRPCMessage, error) {
	msg := &IncomingRPCMessage{
		IncomingRPCMessage: data,
	}
	switch data.BugleRoute {
	case gmproto.BugleRoute_PairEvent:
		msg.Pair = &gmproto.RPCPairData{}
		err := proto.Unmarshal(data.GetMessageData(), msg.Pair)
		if err != nil {
			c.Logger.Trace().
				Str("data", base64.StdEncoding.EncodeToString(msg.GetMessageData())).
				Msg("Errored pair event content")
			return nil, fmt.Errorf("failed to decode pair event: %w", err)
		}
	case gmproto.BugleRoute_GaiaEvent:
		msg.Gaia = &gmproto.RPCGaiaData{}
		err := proto.Unmarshal(data.GetMessageData(), msg.Gaia)
		if err != nil {
			c.Logger.Trace().
				Str("data", base64.StdEncoding.EncodeToString(msg.GetMessageData())).
				Msg("Errored gaia event content")
			return nil, fmt.Errorf("failed to decode gaia event: %w", err)
		}
	case gmproto.BugleRoute_DataEvent:
		msg.Message = &gmproto.RPCMessageData{}
		err := proto.Unmarshal(data.GetMessageData(), msg.Message)
		if err != nil {
			c.Logger.Trace().
				Str("data", base64.StdEncoding.EncodeToString(msg.GetMessageData())).
				Msg("Errored data event content")
			return nil, fmt.Errorf("failed to decode data event: %w", err)
		}
		responseStruct, ok := responseType[msg.Message.GetAction()]
		if ok {
			msg.DecryptedMessage = responseStruct.ProtoReflect().New().Interface()
		}
		if msg.Message.EncryptedData != nil {
			msg.DecryptedData, err = c.AuthData.RequestCrypto.Decrypt(msg.Message.EncryptedData)
			if err != nil {
				return nil, fmt.Errorf("failed to decrypt data event: %w", err)
			}
			if msg.DecryptedMessage != nil {
				err = proto.Unmarshal(msg.DecryptedData, msg.DecryptedMessage)
				if err != nil {
					c.Logger.Trace().
						Str("data", base64.StdEncoding.EncodeToString(msg.DecryptedData)).
						Msg("Errored decrypted data event content")
					return nil, fmt.Errorf("failed to decode decrypted data event: %w", err)
				}
			}
		} else if msg.Message.EncryptedData2 != nil {
			msg.DecryptedData, err = c.AuthData.RequestCrypto.Decrypt(msg.Message.EncryptedData2)
			if err != nil {
				return nil, fmt.Errorf("failed to decrypt field 2 in data event: %w", err)
			}
			var ed2c gmproto.EncryptedData2Container
			err = proto.Unmarshal(msg.DecryptedData, &ed2c)
			if err != nil {
				c.Logger.Trace().
					Str("data", base64.StdEncoding.EncodeToString(msg.DecryptedData)).
					Msg("Errored decrypted data event content")
				return nil, fmt.Errorf("failed to decode decrypted field 2 data event: %w", err)
			}
			// Hacky hack to have User.handleAccountChange do the right-ish thing on startup
			if strings.ContainsRune(ed2c.GetAccountChange().GetAccount(), '@') {
				c.triggerEvent(&events.AccountChange{
					AccountChangeOrSomethingEvent: ed2c.GetAccountChange(),
					IsFake:                        true,
				})
			}
		}
	default:
		return nil, fmt.Errorf("unknown bugle route %d", data.BugleRoute)
	}
	return msg, nil
}

func (c *Client) deduplicateHash(id string, hash [32]byte) bool {
	const recentUpdatesLen = len(c.recentUpdates)
	for i := c.recentUpdatesPtr + recentUpdatesLen - 1; i >= c.recentUpdatesPtr; i-- {
		if c.recentUpdates[i%recentUpdatesLen].id == id {
			if c.recentUpdates[i%recentUpdatesLen].hash == hash {
				return true
			} else {
				break
			}
		}
	}
	c.recentUpdates[c.recentUpdatesPtr] = updateDedupItem{id: id, hash: hash}
	c.recentUpdatesPtr = (c.recentUpdatesPtr + 1) % recentUpdatesLen
	return false
}

func (c *Client) logContent(res *IncomingRPCMessage, thingID string, contentHash []byte) {
	if c.Logger.Trace().Enabled() && (res.DecryptedData != nil || res.DecryptedMessage != nil) {
		evt := c.Logger.Trace().Bool("is_old", res.IsOld)
		if res.DecryptedMessage != nil {
			evt.Str("proto_name", string(res.DecryptedMessage.ProtoReflect().Descriptor().FullName()))
		}
		if res.DecryptedData != nil {
			evt.Str("data", base64.StdEncoding.EncodeToString(res.DecryptedData))
			if contentHash != nil {
				evt.Str("thing_id", thingID)
				evt.Hex("data_hash", contentHash)
			}
		} else {
			evt.Str("data", "<null>")
		}
		evt.Msg("Got event")
	}
}

func (c *Client) deduplicateUpdate(id string, msg *IncomingRPCMessage) bool {
	if msg.DecryptedData != nil {
		contentHash := sha256.Sum256(msg.DecryptedData)
		if c.deduplicateHash(id, contentHash) {
			c.Logger.Trace().
				Str("thing_id", id).
				Hex("data_hash", contentHash[:]).
				Bool("is_old", msg.IsOld).
				Msg("Ignoring duplicate update")
			return true
		}
		c.logContent(msg, id, contentHash[:])
	}
	return false
}

func (c *Client) HandleRPCMsg(rawMsg *gmproto.IncomingRPCMessage) {
	msg, err := c.decryptInternalMessage(rawMsg)
	if err != nil {
		c.Logger.Err(err).Str("message_id", rawMsg.ResponseID).Msg("Failed to decode incoming RPC message")
		c.sessionHandler.queueMessageAck(rawMsg.ResponseID)
		return
	}

	c.sessionHandler.queueMessageAck(msg.ResponseID)
	if c.sessionHandler.receiveResponse(msg) {
		return
	}
	logEvt := c.Logger.Debug().
		Str("message_id", msg.ResponseID).
		Stringer("bugle_route", msg.BugleRoute)
	if msg.Message != nil {
		logEvt.Stringer("message_action", msg.Message.Action)
	}
	logEvt.Msg("Received message")
	switch msg.BugleRoute {
	case gmproto.BugleRoute_PairEvent:
		c.handlePairingEvent(msg)
	case gmproto.BugleRoute_GaiaEvent:
		c.handleGaiaPairingEvent(msg)
	case gmproto.BugleRoute_DataEvent:
		if c.skipCount > 0 {
			c.skipCount--
			msg.IsOld = true
		}
		c.handleUpdatesEvent(msg)
	}
}

type WrappedMessage struct {
	*gmproto.Message
	IsOld bool
	Data  []byte
}

var hackyLoggedOutBytes = []byte{0x72, 0x00}

func (c *Client) handleUpdatesEvent(msg *IncomingRPCMessage) {
	switch msg.Message.Action {
	case gmproto.ActionType_GET_UPDATES:
		if msg.DecryptedData == nil && bytes.Equal(msg.Message.UnencryptedData, hackyLoggedOutBytes) {
			c.triggerEvent(&events.GaiaLoggedOut{})
			return
		}
		if !msg.IsOld {
			c.bumpNextDataReceiveCheck(DefaultBugleDefaultCheckInterval)
		}
		data, ok := msg.DecryptedMessage.(*gmproto.UpdateEvents)
		if !ok {
			c.Logger.Error().
				Type("data_type", msg.DecryptedMessage).
				Bool("is_old", msg.IsOld).
				Msg("Unexpected data type in GET_UPDATES event")
			return
		}

		switch evt := data.Event.(type) {
		case *gmproto.UpdateEvents_UserAlertEvent:
			c.logContent(msg, "", nil)
			if msg.IsOld {
				return
			}
			c.triggerEvent(evt.UserAlertEvent)

		case *gmproto.UpdateEvents_SettingsEvent:
			c.Logger.Debug().
				Str("data", base64.StdEncoding.EncodeToString(msg.DecryptedData)).
				Bool("is_old", msg.IsOld).
				Msg("Got settings event")
			c.triggerEvent(evt.SettingsEvent)

		case *gmproto.UpdateEvents_ConversationEvent:
			for _, part := range evt.ConversationEvent.GetData() {
				if c.deduplicateUpdate(part.GetConversationID(), msg) {
					return
				} else if msg.IsOld {
					c.Logger.Debug().Str("conv_id", part.ConversationID).Msg("Ignoring old conversation event")
					continue
				}
				c.triggerEvent(part)
			}

		case *gmproto.UpdateEvents_MessageEvent:
			for _, part := range evt.MessageEvent.GetData() {
				if c.deduplicateUpdate(part.GetMessageID(), msg) {
					return
				}
				c.triggerEvent(&WrappedMessage{
					Message: part,
					IsOld:   msg.IsOld,
					Data:    msg.DecryptedData,
				})
			}

		case *gmproto.UpdateEvents_TypingEvent:
			c.logContent(msg, "", nil)
			if msg.IsOld {
				return
			}
			c.triggerEvent(evt.TypingEvent.GetData())

		case *gmproto.UpdateEvents_AccountChange:
			c.logContent(msg, "", nil)
			c.triggerEvent(&events.AccountChange{
				AccountChangeOrSomethingEvent: evt.AccountChange,
			})

		default:
			c.Logger.Warn().
				Str("evt_data", base64.StdEncoding.EncodeToString(msg.GetMessageData())).
				Str("decrypted_data", base64.StdEncoding.EncodeToString(msg.DecryptedData)).
				Msg("Got unknown event type")
		}
	default:
		c.Logger.Debug().
			Str("evt_data", base64.StdEncoding.EncodeToString(msg.GetMessageData())).
			Str("request_id", msg.Message.SessionID).
			Str("action_type", msg.Message.Action.String()).
			Bool("is_old", msg.IsOld).
			Msg("Got unexpected response")
	}
}
