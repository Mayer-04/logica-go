CREATE TABLE authors (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL CHECK (char_length(trim(name)) > 0),
  bio  text
);

-- name: GetAuthor :one
SELECT id, name, bio
FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT *
FROM authors
ORDER BY name ASC;

-- name: CreateAuthor :one
INSERT INTO authors (name, bio)
VALUES ($1, $2)
RETURNING id, name, bio;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;