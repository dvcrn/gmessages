INSERT INTO "user" (bridge_id, mxid, management_room, access_token)
SELECT '', mxid, management_room, access_token
FROM user_old;

CREATE TABLE gmessages_login_prefix(
    -- only: postgres
    prefix BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    -- only: sqlite (line commented)
--	prefix INTEGER PRIMARY KEY,
    login_id TEXT NOT NULL,

    CONSTRAINT gmessages_login_prefix_login_id_key UNIQUE (login_id)
);

CREATE TABLE gmessages_version (version INTEGER, compat INTEGER);
INSERT INTO gmessages_version (version, compat) VALUES (1, 1);

INSERT INTO gmessages_login_prefix (prefix, login_id)
SELECT rowid, COALESCE(phone_id, CAST(rowid AS TEXT))
FROM user_old;

-- only: postgres
SELECT setval('gmessages_login_prefix_prefix_seq', (SELECT MAX(prefix)+1 FROM gmessages_login_prefix), FALSE);

INSERT INTO user_login (bridge_id, user_mxid, id, remote_name, space_room, metadata)
SELECT
    '', -- bridge_id
    mxid, -- user_mxid
    phone_id, -- id
    '', -- remote_name
    space_room, -- space_room
    -- only: postgres
    jsonb_build_object
-- only: sqlite (line commented)
--  json_object
    (
        'session', json(session),
        'id_prefix', CAST(rowid AS TEXT),
        'self_participant_ids', json(self_participant_ids),
        'sim_metadata', json(sim_metadata),
        'settings', json(settings)
    ) -- metadata
FROM user_old WHERE phone_id<>'';

INSERT INTO ghost (
    bridge_id, id, name, avatar_id, avatar_hash, avatar_mxc,
    name_set, avatar_set, contact_info_set, is_bot, identifiers, metadata
)
SELECT
    '', -- bridge_id
    (CAST(receiver AS TEXT) || '.' || id), -- id
    name, -- name
    CASE WHEN avatar_hash IS NULL THEN '' ELSE
        -- only: postgres
        'hash:' || encode(avatar_hash, 'hex')
        -- only: sqlite (line commented)
--      'hash:' || hex(avatar_hash)
    END, -- avatar_id
    CASE WHEN avatar_hash IS NULL THEN '' ELSE
        -- only: postgres
        encode(avatar_hash, 'hex')
        -- only: sqlite (line commented)
--      hex(avatar_hash)
    END, -- avatar_hash
    avatar_mxc,
    name_set,
    avatar_set,
    contact_info_set,
    false, -- is_bot
    '[]', -- identifiers
    -- only: postgres
    jsonb_build_object
-- only: sqlite (line commented)
--  json_object
    (
        'contact_id', contact_id,
        'phone', phone,
        'avatar_update_ts', avatar_update_ts
    ) -- metadata
FROM puppet_old;

UPDATE ghost SET avatar_id='', avatar_hash='' WHERE avatar_hash='0000000000000000000000000000000000000000000000000000000000000000';

INSERT INTO portal (
    bridge_id, id, receiver, mxid, other_user_id, name, topic, avatar_id, avatar_hash, avatar_mxc,
    name_set, avatar_set, topic_set, in_space, room_type, metadata
)
SELECT
    '', -- bridge_id
    (CAST(receiver AS TEXT) || '.' || id), -- id
    (SELECT login_id FROM gmessages_login_prefix WHERE prefix=portal_old.receiver), -- receiver
    mxid,
    CASE WHEN other_user IS NOT NULL THEN (CAST(receiver AS TEXT) || '.' || other_user) END, -- other_user_id
    name,
    '', -- topic
    '', -- avatar_id
    '', -- avatar_hash
    '', -- avatar_mxc
    name_set,
    false, -- avatar_set
    false, -- topic_set
    false, -- in_space (spaceness is stored in user_portal)
    CASE WHEN other_user IS NOT NULL THEN 'dm' ELSE '' END, -- room_type
    -- only: postgres
    jsonb_build_object
-- only: sqlite (line commented)
--  json_object
    (
        'type', type,
        'send_mode', send_mode,
        'force_rcs', force_rcs
    ) -- metadata
FROM portal_old;
-- only: sqlite
UPDATE portal SET metadata=replace(replace(metadata, '"force_rcs":1', '"force_rcs":true'), '"force_rcs":0', '"force_rcs":false');

INSERT INTO user_portal (
    bridge_id, user_mxid, login_id, portal_id, portal_receiver, in_space, preferred
)
SELECT
    '', -- bridge_id
    (SELECT mxid FROM user_old WHERE rowid=receiver), -- user_mxid
    (SELECT login_id FROM gmessages_login_prefix WHERE prefix=portal_old.receiver), -- login_id
    (CAST(receiver AS TEXT) || '.' || id), -- portal_id
    (SELECT login_id FROM gmessages_login_prefix WHERE prefix=portal_old.receiver), -- portal_receiver
    in_space,
    false
FROM portal_old;

INSERT INTO ghost (
    bridge_id, id, name, avatar_id, avatar_hash, avatar_mxc,
    name_set, avatar_set, contact_info_set, is_bot, identifiers, metadata
)
SELECT DISTINCT
    '', CAST(conv_receiver AS TEXT) || '.' || sender, '', '', '', '', false, false, false, false,
    -- only: postgres
    '[]'::jsonb,
    -- only: sqlite (line commented)
--   '[]',
    '{}'
FROM message_old
WHERE true
ON CONFLICT DO NOTHING;

INSERT INTO message (
    bridge_id, id, part_id, mxid, room_id, room_receiver, sender_id, sender_mxid,
    timestamp, edit_count, metadata
)
SELECT
    '', -- bridge_id
    (CAST(conv_receiver AS TEXT) || '.' || id), -- id
    '', -- part_id,
    mxid,
    (CAST(conv_receiver AS TEXT) || '.' || conv_id), -- room_id
    (SELECT login_id FROM gmessages_login_prefix WHERE prefix=conv_receiver), -- room_receiver
    (CAST(conv_receiver AS TEXT) || '.' || sender), -- sender_id
    '', -- sender_mxid
    timestamp * 1000,
    0, -- edit_count
    status -- metadata
FROM message_old;
-- TODO split out parts from status?

INSERT INTO reaction (
    bridge_id, message_id, message_part_id, sender_id, emoji_id,
    room_id, room_receiver, mxid, timestamp, emoji, metadata
)
SELECT
    '', -- bridge_id
    (CAST(conv_receiver AS TEXT) || '.' || msg_id), -- message_id
    '', -- message_part_id
    (CAST(conv_receiver AS TEXT) || '.' || sender), -- sender_id
    '', -- emoji_id
    (CAST(conv_receiver AS TEXT) || '.' || conv_id), -- room_id
    (SELECT login_id FROM gmessages_login_prefix WHERE prefix=conv_receiver), -- room_receiver
    mxid,
    (SELECT (timestamp * 1000) + 1 FROM message_old WHERE conv_receiver=reaction_old.conv_receiver and id=reaction_old.msg_id), -- timestamp
    reaction, -- emoji
    '{}' -- metadata
FROM reaction_old
