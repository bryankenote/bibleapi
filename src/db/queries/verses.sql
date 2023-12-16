-- name: GetVerse :one
SELECT * FROM verses
WHERE translation = ?
AND book = ?
AND chapter = ?
AND verse = ?
LIMIT 1;

-- name: GetChapter :many
SELECT * FROM verses
WHERE translation = ?
AND book = ?
AND chapter = ?;

-- name: GetBook :many
SELECT * FROM verses
WHERE translation = ?
AND book = ?;

-- name: CreateVerse :one
INSERT INTO verses (
	translation, book, chapter, verse, content
) VALUES (
	?, ?, ?, ?, ?
)
RETURNING *;
