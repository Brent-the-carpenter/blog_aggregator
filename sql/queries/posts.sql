-- name: CreatePost :one
INSERT INTO posts(
	id ,
	created_at,
	updated_at,
	title,
	url,
	description,
	published_at,
	feed_id
)VALUES(
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8
	)
RETURNING *;


-- name: GetPosts :many
SELECT posts.* , feeds.name as feed_name FROM posts JOIN feeds ON posts.feed_id = feeds.id ORDER BY published_at DESC , feeds.name LIMIT $1;
