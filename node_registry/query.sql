-- name: AddNode :exec
INSERT INTO node(label, base_url, is_online)
    VALUES (?, ?, ?)
    ON CONFLICT (label) DO UPDATE set base_url = excluded.base_url, is_online = excluded.is_online;

-- name: ListNodes :many
SELECT * FROM node;

-- name: ListOnlineNodes :many
SELECT * FROM node WHERE is_online = 1;