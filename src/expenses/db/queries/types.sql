-- name: GetTypes :many
SELECT * FROM "expenseTypes";

-- name: GetType :one
SELECT * FROM "expenseTypes" WHERE "TypeID" = ?;

-- name: InsertType :execresult
INSERT INTO "expenseTypes"("TypeName") VALUES(?);

-- name: UpdateType :execresult
UPDATE expenseTypes SET "TypeName" = ? WHERE "TypeID" = ?;

-- name: DeleteType :execresult
DELETE FROM "expenseTypes" WHERE "TypeID" = ?;
