-- name: GetExpenses :many
SELECT 
  sqlc.embed(expenses),
  sqlc.embed(stores),
  sqlc.embed(categories),
  sqlc.embed(users),
  sqlc.embed(types)
FROM expenses
JOIN 
  Stores ON stores.StoreID = expenses.StoreID
JOIN 
  Categories ON categories.CategoryID = expenses.CategoryID
JOIN 
  Users ON UserID = OwnerUserId
JOIN
  "expenseTypes" as types ON types.TypeID = expenses.TypeID
WHERE
  (sqlc.narg(startDate) IS NULL OR expenses."ExpDate" >= sqlc.narg(startDate))
  AND
  (sqlc.narg(endDate) IS NULL OR expenses."ExpDate" <= sqlc.narg(endDate));

-- name: GetExpense :one
SELECT 
  sqlc.embed(expenses),
  sqlc.embed(stores),
  sqlc.embed(categories),
  sqlc.embed(users),
  sqlc.embed(types)
FROM expenses
JOIN 
  Stores ON stores.StoreID = expenses.StoreID
JOIN 
  Categories ON categories.CategoryID = expenses.CategoryID
JOIN 
  Users ON "UserID" = "OwnerUserId"
JOIN
  "expenseTypes" as types ON types.TypeID = expenses.TypeID
WHERE 
  "ExpID" = ?;

-- name: InsertExpense :execresult
INSERT INTO expenses(
  "Description","Value",
  "StoreID",
  "CategoryID",
  "TypeID",
  "OwnerUserID",
  "ExpDate","CreationDate"
)
VALUES(?, ?, ? , ?, ?, ?, ?, ?);

-- name: DeleteExpense :execresult
DELETE FROM expenses WHERE "ExpID" = ?;

-- name: UpdateExpense :execresult
UPDATE expenses
SET
  "Description" = ?,
  "Value" = ?,
  "StoreID" = ?,
  "CategoryID" = ?,
  "TypeID" = ?,
  "OwnerUserID" = ?,
  "ExpDate" = ?
WHERE "ExpID" = ?;
