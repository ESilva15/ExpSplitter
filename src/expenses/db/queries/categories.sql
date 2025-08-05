-- name: GetCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE "CategoryID" = ?;

-- name: UpdateCategory :execresult
UPDATE categories SET "CategoryName" = ? WHERE "CategoryID" = ?;

-- name: DeleteCategory :execresult
DELETE FROM categories WHERE "CategoryID" = ?;

-- name: InsertCategory :execresult
INSERT INTO categories("CategoryName") VALUES(?)
