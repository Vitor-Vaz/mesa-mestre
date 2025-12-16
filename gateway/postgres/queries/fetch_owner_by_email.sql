-- name: FetchOwnerByEmail :one
SELECT id, name, email, created_at
FROM owners
WHERE email = $1;