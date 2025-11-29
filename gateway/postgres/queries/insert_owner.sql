-- name: InsertOwner :exec
INSERT INTO owners (name, email)
VALUES ($1, $2);