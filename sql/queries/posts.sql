-- name: CreatePost :one
INSERT INTO posts (
  id, 
  created_at, 
  updated_at, 
  title, 
  url, 
  description, 
  published_at, 
  feed_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
--

-- name: GetPostsForUser :many
SELECT t1.*, t3.name as feed_name FROM posts t1
INNER JOIN feed_follows t2
  ON t1.feed_id = t2.feed_id
INNER JOIN feeds t3
  ON t1.feed_id = t3.id
WHERE t2.user_id = $1
ORDER BY t1.published_at DESC 
LIMIT $2;
--
