-- name: AddNode :exec
INSERT INTO node (label, base_url, is_online)
    VALUES (?, ?, ?);