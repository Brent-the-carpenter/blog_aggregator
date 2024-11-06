-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(
        id,
        created_at,
        updated_at,
        user_id,
        feed_id
    )
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_feed_follow.*, 
    feeds.name AS feed_name,
    users.name AS user_name 
FROM inserted_feed_follow 
JOIN users ON users.id = inserted_feed_follow.user_id 
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
WHERE inserted_feed_follow.id = inserted_feed_follow.id;

-- name: GetFeedFollowsForUser :many
SELECT users.name as userName , feeds.name as feedName FROM feed_follows
JOIN users ON feed_follows.user_id = users.id 
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1 ;

-- name: UnfollowFeed :exec
 
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;
