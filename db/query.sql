-- name: GetAllHouses :many
SELECT * FROM houses;

-- name: GetStudentsByHouse :many
SELECT * FROM students WHERE house_id = $1;

-- name: CreateHouse :one
INSERT INTO houses (id, first_name, last_name, head_of_house) 
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateStudent :one
INSERT INTO students (id, first_name, last_name, age, house_id) 
VALUES ($1, $2, $3, $4, $5) RETURNING *;