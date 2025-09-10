-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE "UserID" = $1;

-- name: GetUserByName :one
SELECT * FROM users WHERE "UserName" = $1;
