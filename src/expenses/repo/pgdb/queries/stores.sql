-- name: GetStores :many
SELECT * FROM stores;

-- name: GetStore :one
SELECT * FROM stores WHERE "StoreID" = $1;

-- name: InsertStore :execresult
INSERT INTO stores("StoreName") VALUES($1);

-- name: UpdateStore :execresult
UPDATE stores SET "StoreName" = $1 WHERE "StoreID" = $2;

-- name: DeleteStore :execresult
DELETE FROM stores WHERE "StoreID" = $1;
