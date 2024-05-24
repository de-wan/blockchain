-- name: AddNode :exec
INSERT INTO node(label, base_url, is_online)
    VALUES (?, ?, ?)
    ON CONFLICT (label) DO UPDATE set base_url = excluded.base_url, is_online = excluded.is_online;