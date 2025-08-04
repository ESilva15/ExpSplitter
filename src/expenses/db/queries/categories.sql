-- name: GetCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE "CategoryID" = ?;
