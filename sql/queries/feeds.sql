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
