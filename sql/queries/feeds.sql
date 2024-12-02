-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedSummaries :many
SELECT 
  t1.name as feed_name,
  t1.url,
  t2.name as user_name
FROM feeds t1
INNER JOIN users t2
on t1.user_id = t2.id;

-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT
  t1.*,
  t2.name as feed_name,
  t3.name as user_name
FROM insert_feed_follow t1
INNER JOIN feeds t2 ON t1.feed_id = t2.id
INNER JOIN users t3 ON t1.user_id = t3.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT
  t1.*,
  t2.name as feed_name,
  t3.name as user_name
FROM feed_follows t1
INNER JOIN feeds t2 ON t1.feed_id = t2.id
INNER JOIN users t3 ON t1.user_id = t3.id
WHERE t3.name = $1;
