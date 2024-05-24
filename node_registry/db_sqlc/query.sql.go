// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db_sqlc

import (
	"context"
)

const addNode = `-- name: AddNode :exec
INSERT INTO node(label, base_url, is_online)
    VALUES (?, ?, ?)
    ON CONFLICT (label) DO UPDATE set base_url = excluded.base_url, is_online = excluded.is_online
`

type AddNodeParams struct {
	Label    string
	BaseUrl  string
	IsOnline bool
}

func (q *Queries) AddNode(ctx context.Context, arg AddNodeParams) error {
	_, err := q.db.ExecContext(ctx, addNode, arg.Label, arg.BaseUrl, arg.IsOnline)
	return err
}

const listNodes = `-- name: ListNodes :many
SELECT id, label, base_url, is_online, last_heartbeat, created_at FROM node
`

func (q *Queries) ListNodes(ctx context.Context) ([]Node, error) {
	rows, err := q.db.QueryContext(ctx, listNodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.ID,
			&i.Label,
			&i.BaseUrl,
			&i.IsOnline,
			&i.LastHeartbeat,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOnlineNodes = `-- name: ListOnlineNodes :many
SELECT id, label, base_url, is_online, last_heartbeat, created_at FROM node WHERE is_online = 1
`

func (q *Queries) ListOnlineNodes(ctx context.Context) ([]Node, error) {
	rows, err := q.db.QueryContext(ctx, listOnlineNodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.ID,
			&i.Label,
			&i.BaseUrl,
			&i.IsOnline,
			&i.LastHeartbeat,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
