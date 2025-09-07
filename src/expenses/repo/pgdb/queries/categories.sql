-- name: GetCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE "CategoryID" = $1;

-- name: UpdateCategory :execresult
UPDATE categories SET "CategoryName" = $1 WHERE "CategoryID" = $2;

-- name: DeleteCategory :execresult
DELETE FROM categories WHERE "CategoryID" = $1;

-- name: InsertCategory :execresult
INSERT INTO categories("CategoryName") VALUES($1);
