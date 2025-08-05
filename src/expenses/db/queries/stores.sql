-- name: GetStores :many
SELECT * FROM stores;

-- name: GetStore :one
SELECT * FROM stores WHERE "StoreID" = ?;

-- name: InsertStore :execresult
INSERT INTO stores("StoreName") VALUES(?);

-- name: UpdateStore :execresult
UPDATE stores SET "StoreName" = ? WHERE "StoreID" = ?;

-- name: DeleteStore :execresult
DELETE FROM stores WHERE "StoreID" = ?;
