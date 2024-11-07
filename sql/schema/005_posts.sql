-- +goose Up
CREATE TABLE posts(
	id UUID PRIMARY KEY NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL ,
	title VARCHAR(255) NOT NULL,
	url  VARCHAR(255) UNIQUE NOT NULL,
	description TEXT ,
	published_at TIMESTAMP NOT NULL,
	feed_id UUID NOT NULL ,
	CONSTRAINT fk_feed_id FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE posts;
