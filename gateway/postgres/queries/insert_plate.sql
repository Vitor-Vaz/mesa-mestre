-- name: InsertPlate :exec
INSERT INTO plates (plate_name, plate_description, price)
VALUES ($1, $2, $3);