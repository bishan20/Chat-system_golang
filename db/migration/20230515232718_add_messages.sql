-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL
);

ALTER TABLE messages ADD FOREIGN KEY (sender_id) REFERENCES users(id);

ALTER TABLE messages ADD FOREIGN KEY (receiver_id) REFERENCES users(id);

ALTER TABLE messages ADD COLUMN IF NOT EXISTS is_delivered bool NOT NULL DEFAULT FALSE;

ALTER TABLE messages ADD COLUMN IF NOT EXISTS sent_at timestamp NOT NULL DEFAULT now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
