-- name: StoreMessage :one
INSERT INTO messages (
    message,
    sender_id,
    receiver_id
) VALUES (
    $1, $2, $3
)  RETURNING *; 

-- name: ListMessages :many
SELECT * FROM messages
WHERE (sender_id = $1 AND receiver_id = $2)
OR (sender_id = $2 and receiver_id = $1);

-- name: UpdateMessage :one
UPDATE messages
SET message = $3
WHERE id = $1 AND sender_id = $2
RETURNING *;

-- name: UpdateMessageDelivery :one
UPDATE messages
SET is_delivered = TRUE
WHERE id = $1
RETURNING id;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = $1 AND sender_id = $2;