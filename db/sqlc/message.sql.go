// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: message.sql

package db

import (
	"context"
)

const deleteMessage = `-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = $1 AND sender_id = $2
`

type DeleteMessageParams struct {
	ID       int32 `db:"id" json:"id"`
	SenderID int32 `db:"sender_id" json:"sender_id"`
}

func (q *Queries) DeleteMessage(ctx context.Context, arg DeleteMessageParams) error {
	_, err := q.exec(ctx, q.deleteMessageStmt, deleteMessage, arg.ID, arg.SenderID)
	return err
}

const listMessages = `-- name: ListMessages :many
SELECT id, message, sender_id, receiver_id, is_delivered, sent_at FROM messages
WHERE (sender_id = $1 AND receiver_id = $2)
OR (sender_id = $2 and receiver_id = $1)
`

type ListMessagesParams struct {
	SenderID   int32 `db:"sender_id" json:"sender_id"`
	ReceiverID int32 `db:"receiver_id" json:"receiver_id"`
}

func (q *Queries) ListMessages(ctx context.Context, arg ListMessagesParams) ([]Message, error) {
	rows, err := q.query(ctx, q.listMessagesStmt, listMessages, arg.SenderID, arg.ReceiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Message,
			&i.SenderID,
			&i.ReceiverID,
			&i.IsDelivered,
			&i.SentAt,
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

const storeMessage = `-- name: StoreMessage :one
INSERT INTO messages (
    message,
    sender_id,
    receiver_id
) VALUES (
    $1, $2, $3
)  RETURNING id, message, sender_id, receiver_id, is_delivered, sent_at
`

type StoreMessageParams struct {
	Message    string `db:"message" json:"message"`
	SenderID   int32  `db:"sender_id" json:"sender_id"`
	ReceiverID int32  `db:"receiver_id" json:"receiver_id"`
}

func (q *Queries) StoreMessage(ctx context.Context, arg StoreMessageParams) (Message, error) {
	row := q.queryRow(ctx, q.storeMessageStmt, storeMessage, arg.Message, arg.SenderID, arg.ReceiverID)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.SenderID,
		&i.ReceiverID,
		&i.IsDelivered,
		&i.SentAt,
	)
	return i, err
}

const updateMessage = `-- name: UpdateMessage :one
UPDATE messages
SET message = $3
WHERE id = $1 AND sender_id = $2
RETURNING id, message, sender_id, receiver_id, is_delivered, sent_at
`

type UpdateMessageParams struct {
	ID       int32  `db:"id" json:"id"`
	SenderID int32  `db:"sender_id" json:"sender_id"`
	Message  string `db:"message" json:"message"`
}

func (q *Queries) UpdateMessage(ctx context.Context, arg UpdateMessageParams) (Message, error) {
	row := q.queryRow(ctx, q.updateMessageStmt, updateMessage, arg.ID, arg.SenderID, arg.Message)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.SenderID,
		&i.ReceiverID,
		&i.IsDelivered,
		&i.SentAt,
	)
	return i, err
}

const updateMessageDelivery = `-- name: UpdateMessageDelivery :one
UPDATE messages
SET is_delivered = TRUE
WHERE id = $1
RETURNING id
`

func (q *Queries) UpdateMessageDelivery(ctx context.Context, id int32) (int32, error) {
	row := q.queryRow(ctx, q.updateMessageDeliveryStmt, updateMessageDelivery, id)
	err := row.Scan(&id)
	return id, err
}
