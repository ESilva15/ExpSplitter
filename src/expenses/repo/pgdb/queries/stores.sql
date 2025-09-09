-- name: GetStores :many
SELECT * FROM stores;

-- name: GetStore :one
SELECT * FROM stores WHERE "StoreID" = $1;

-- name: InsertStore :execresult
INSERT INTO stores("StoreName", "NIF") VALUES($1, $2);

-- name: UpdateStore :execresult
UPDATE stores SET "StoreName" = $1, "NIF" = $2 WHERE "StoreID" = $3;

-- name: DeleteStore :execresult
DELETE FROM stores WHERE "StoreID" = $1;

-- name: GetStoreByNIF :one
SELECT * FROM stores WHERE "NIF" = $1;
