-- +goose Up
ALTER TABLE feeds 
ADD last_fetched_at TIMESTAMP  NULL ;

CREATE INDEX idx_last_fetched ON feeds(last_fetched_at);

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_fetched_at ;

DROP INDEX IF EXISTS idx_last_fetched;
