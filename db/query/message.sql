-- name: StoreMessage :one
INSERT INTO messages (
    message,
    sender_id,
    receiver_id
) VALUES (
    $1, $2, $3
)  RETURNING *; 

-- name: UpdateMessage :one
UPDATE messages
SET message = $3
WHERE id = $1 AND sender_id = $2
RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = $1 AND sender_id = $2;