// mautrix-gmessages - A Matrix-Google Messages puppeting bridge.
// Copyright (C) 2023 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"go.mau.fi/util/random"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"

	"go.mau.fi/mautrix-gmessages/database"
)

func (portal *Portal) initialForwardBackfill(user *User, markRead, allowNotify bool) {
	// This is only called from CreateMatrixRoom which locks forwardBackfillLock
	defer portal.forwardBackfillLock.Unlock()
	log := portal.zlog.With().
		Str("action", "initial forward backfill").
		Logger()
	ctx := log.WithContext(context.TODO())

	portal.forwardBackfill(ctx, user, time.Time{}, portal.bridge.Config.Bridge.Backfill.InitialLimit, markRead, allowNotify)
}

const recentBackfillDelay = 15 * time.Second

type pendingBackfill struct {
	cancel        context.CancelFunc
	lastMessageID string
	lastMessageTS time.Time
}

func (portal *Portal) missedForwardBackfill(ctx context.Context, user *User, lastMessageTS time.Time, lastMessageID string, markRead, markReadIfNoBackfill bool) {
	if portal.bridge.Config.Bridge.Backfill.MissedLimit == 0 {
		if markRead && markReadIfNoBackfill {
			user.markSelfReadFull(ctx, portal, lastMessageID)
		}
		return
	}
	log := portal.zlog.With().
		Str("action", "missed forward backfill").
		Str("latest_message_id", lastMessageID).
		Logger()
	ctx = log.WithContext(ctx)
	if portal.hasSyncedThisRun && !lastMessageTS.IsZero() && time.Since(lastMessageTS) < 5*time.Minute && portal.lastMessageTS.Before(lastMessageTS) {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		prev := portal.pendingRecentBackfill.Swap(&pendingBackfill{cancel: cancel, lastMessageID: lastMessageID, lastMessageTS: lastMessageTS})
		if prev != nil {
			prev.cancel()
		}
		log.Debug().Msg("Delaying missed forward backfill as latest message is new")
		select {
		case <-time.After(recentBackfillDelay):
		case <-ctx.Done():
			log.Debug().Msg("Backfill was cancelled by a newer backfill")
			return
		}
	}

	portal.forwardBackfillLock.Lock()
	defer portal.forwardBackfillLock.Unlock()
	portal.hasSyncedThisRun = true
	if !lastMessageTS.IsZero() {
		if portal.lastMessageTS.IsZero() {
			lastMsg, err := portal.bridge.DB.Message.GetLastInChat(ctx, portal.Key)
			if err != nil {
				log.Err(err).Msg("Failed to get last message in chat")
				return
			} else if lastMsg == nil {
				log.Debug().Msg("No messages in chat")
			} else {
				portal.lastMessageTS = lastMsg.Timestamp
			}
		}
		if !lastMessageTS.After(portal.lastMessageTS) {
			log.Trace().
				Time("latest_message_ts", lastMessageTS).
				Str("latest_message_id", lastMessageID).
				Time("last_bridged_ts", portal.lastMessageTS).
				Msg("Nothing to backfill")
			if markRead && markReadIfNoBackfill {
				user.markSelfReadFull(ctx, portal, lastMessageID)
			}
			return
		}
	}
	log.Info().
		Time("latest_message_ts", lastMessageTS).
		Str("latest_message_id", lastMessageID).
		Time("last_bridged_ts", portal.lastMessageTS).
		Msg("Backfilling missed messages")
	portal.forwardBackfill(ctx, user, portal.lastMessageTS, portal.bridge.Config.Bridge.Backfill.MissedLimit, markRead, true)
}

func (portal *Portal) deterministicEventID(messageID string, part int) id.EventID {
	data := fmt.Sprintf("%s/gmessages/%s/%d", portal.MXID, messageID, part)
	sum := sha256.Sum256([]byte(data))
	return id.EventID(fmt.Sprintf("$%s:messages.google.com", base64.RawURLEncoding.EncodeToString(sum[:])))
}

func (portal *Portal) forwardBackfill(ctx context.Context, user *User, after time.Time, limit int, markRead, allowNotify bool) bool {
	if limit == 0 {
		return false
	}

	log := zerolog.Ctx(ctx)
	// TODO this should cancel if the context is canceled
	resp, err := user.Client.FetchMessages(portal.ID, int64(limit), nil)
	if err != nil {
		portal.zlog.Error().Err(err).Msg("Failed to fetch messages")
		return false
	}
	log.Debug().
		Int64("total_messages", resp.TotalMessages).
		Int("message_count", len(resp.Messages)).
		Msg("Got message chunk to backfill")

	batchSending := portal.bridge.SpecVersions.Supports(mautrix.BeeperFeatureBatchSending)
	converted := make([]*ConvertedMessage, 0, len(resp.Messages))
	maxTS := time.Time{}
	for i := len(resp.Messages) - 1; i >= 0; i-- {
		evt := resp.Messages[i]
		isTooOld := !time.UnixMicro(evt.Timestamp).After(after)
		if portal.handleExistingMessage(ctx, user, evt, isTooOld, nil) || isTooOld {
			continue
		}
		c := portal.convertGoogleMessage(ctx, user, evt, batchSending, nil)
		if c == nil {
			continue
		}
		converted = append(converted, c)
		if c.Timestamp.After(maxTS) {
			maxTS = c.Timestamp
		}
	}
	if len(converted) == 0 {
		log.Debug().Msg("Didn't get any converted messages")
		return false
	}
	log.Debug().
		Int("converted_count", len(converted)).
		Msg("Converted messages for backfill")

	if batchSending {
		var markReadBy id.UserID
		if markRead {
			markReadBy = user.MXID
		}
		allowNotify = allowNotify && time.Since(maxTS) < 24*time.Hour
		portal.backfillSendBatch(ctx, converted, markReadBy, allowNotify)
	} else {
		lastEventID := portal.backfillSendLegacy(ctx, converted)
		if markRead && user.DoublePuppetIntent != nil {
			err = user.DoublePuppetIntent.MarkRead(ctx, portal.MXID, lastEventID)
			if err != nil {
				log.Err(err).Msg("Failed to mark room as read after backfill")
			}
		}
	}
	if maxTS.After(portal.lastMessageTS) {
		portal.lastMessageTS = maxTS
	}
	return true
}

func (portal *Portal) backfillSendBatch(ctx context.Context, converted []*ConvertedMessage, markReadBy id.UserID, allowNotify bool) {
	log := zerolog.Ctx(ctx)
	events := make([]*event.Event, 0, len(converted))
	dbMessages := make([]*database.Message, 0, len(converted))
	for _, msg := range converted {
		dbm := portal.bridge.DB.Message.New()
		dbm.Chat = portal.Key
		dbm.RoomID = portal.MXID
		dbm.ID = msg.ID
		dbm.Sender = msg.SenderID
		dbm.Timestamp = msg.Timestamp
		dbm.Status.Type = msg.Status
		dbm.Status.PartCount = msg.PartCount
		dbm.Status.MediaStatus = msg.MediaStatus
		dbm.Status.MediaParts = make(map[string]database.MediaPart, len(msg.Parts))
		if msg.DontBridge {
			dbm.MXID = id.EventID(fmt.Sprintf("$fake::%s", random.String(37)))
			dbMessages = append(dbMessages, dbm)
			continue
		}

		for i, part := range msg.Parts {
			content := event.Content{
				Parsed: part.Content,
				Raw:    part.Extra,
			}
			eventType := event.EventMessage
			var err error
			eventType, err = portal.encrypt(ctx, msg.Intent, &content, eventType)
			if err != nil {
				log.Err(err).Str("message_id", msg.ID).Int("part", i).Msg("Failed to encrypt event")
				continue
			}
			msg.Intent.AddDoublePuppetValue(&content)
			evt := &event.Event{
				Sender:    msg.Intent.UserID,
				Type:      eventType,
				Timestamp: msg.Timestamp.UnixMilli(),
				ID:        portal.deterministicEventID(msg.ID, i),
				RoomID:    portal.MXID,
				Content:   content,
			}
			events = append(events, evt)
			if dbm.MXID == "" {
				dbm.MXID = evt.ID
				if part.PendingMedia {
					dbm.Status.MediaParts[""] = database.MediaPart{PendingMedia: true}
				}
			} else {
				dbm.Status.MediaParts[part.ID] = database.MediaPart{
					EventID:      evt.ID,
					PendingMedia: part.PendingMedia,
				}
			}
		}
		if dbm.MXID != "" {
			dbMessages = append(dbMessages, dbm)
		}
	}
	if len(dbMessages) == 0 {
		log.Warn().
			Int("converted_count", len(converted)).
			Int("event_count", len(events)).
			Int("db_count", len(dbMessages)).
			Msg("Didn't get any parts to send from converted messages")
		return
	}
	forward := true
	log.Debug().
		Int("event_count", len(events)).
		Bool("forward", forward).
		Bool("mark_read", markReadBy != "").
		Bool("notify", markReadBy == "").
		Msg("Sending batch of messages")
	_, err := portal.MainIntent().BeeperBatchSend(ctx, portal.MXID, &mautrix.ReqBeeperBatchSend{
		Forward:          forward,
		MarkReadBy:       markReadBy,
		SendNotification: forward && markReadBy == "" && allowNotify,
		Events:           events,
	})
	if err != nil {
		log.Err(err).Msg("Failed to send batch of messages")
		return
	}
	err = portal.bridge.DB.Message.MassInsert(ctx, dbMessages)
	if err != nil {
		log.Err(err).Msg("Failed to insert messages to database")
	}
}

func (portal *Portal) backfillSendLegacy(ctx context.Context, converted []*ConvertedMessage) id.EventID {
	var lastEventID id.EventID
	eventIDs := make(map[string]id.EventID)
	for _, msg := range converted {
		msgEventIDs := portal.sendMessageParts(ctx, msg, eventIDs)
		if len(msgEventIDs) > 0 {
			eventIDs[msg.ID] = msgEventIDs[0]
			lastEventID = msgEventIDs[len(msgEventIDs)-1]
		}
	}
	return lastEventID
}
