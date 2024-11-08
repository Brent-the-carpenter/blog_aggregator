-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES(
	$1,
	$2,
	$3,
	$4
	)
	RETURNING *;

-- name: GetUser :one 
Select * FROM users 
where users.name = $1;

-- name: GetUserById :one
Select * FROM users Where users.id = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;
