-- name: GetTypes :many
SELECT * FROM "expenseTypes";

-- name: GetType :one
SELECT * FROM "expenseTypes" WHERE "TypeID" = $1;

-- name: InsertType :execresult
INSERT INTO "expenseTypes"("TypeName") VALUES($1);

-- name: UpdateType :execresult
UPDATE "expenseTypes" SET "TypeName" = $1 WHERE "TypeID" = $2;

-- name: DeleteType :execresult
DELETE FROM "expenseTypes" WHERE "TypeID" = $1;
