-- name: GetExpenses :many
SELECT 
  sqlc.embed(expenses),
  sqlc.embed(stores),
  sqlc.embed(categories),
  sqlc.embed(users),
  sqlc.embed(types)
FROM expenses
JOIN 
  stores ON stores."StoreID" = expenses."StoreID"
JOIN 
  categories ON categories."CategoryID" = expenses."CategoryID"
JOIN 
  users ON users."UserID" = expenses."OwnerUserID"
JOIN
  "expenseTypes" as types ON types."TypeID" = expenses."TypeID"
WHERE
  (sqlc.narg(StartDate)::timestamp IS NULL OR expenses."ExpDate" >= sqlc.narg(StartDate)::timestamp)
  AND
  (sqlc.narg(EndDate)::timestamp IS NULL OR expenses."ExpDate" <= sqlc.narg(EndDate)::timestamp);

-- name: GetExpense :one
SELECT 
  sqlc.embed(expenses),
  sqlc.embed(stores),
  sqlc.embed(categories),
  sqlc.embed(users),
  sqlc.embed(types)
FROM expenses
JOIN 
  stores ON stores."StoreID" = expenses."StoreID"
JOIN 
  categories ON categories."CategoryID" = expenses."CategoryID"
JOIN 
  users ON "UserID" = "OwnerUserID"
JOIN
  "expenseTypes" as types ON types."TypeID" = expenses."TypeID"
WHERE 
  expenses."ExpID" = $1;

-- name: InsertExpense :one
INSERT INTO expenses(
  "Description",
  "Value",
  "StoreID",
  "CategoryID",
  "TypeID",
  "OwnerUserID",
  "ExpDate",
  "PaidOff",
  "SharesEven",
  "CreationDate"
)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING "ExpID";

-- name: DeleteExpense :execresult
WITH delPayments AS (
  DELETE FROM "expensesPayments" WHERE "ExpID" = $1
),
delShares AS (
  DELETE FROM "expensesShares" WHERE "ExpID" = $1
)
DELETE FROM expenses WHERE expenses."ExpID" = $1;

-- name: UpdateExpense :execresult
UPDATE expenses
SET
  "Description" = $1,
  "Value" = $2,
  "StoreID" = $3,
  "CategoryID" = $4,
  "TypeID" = $5,
  "OwnerUserID" = $6,
  "PaidOff" = $7,
  "SharesEven" = $8,
  "ExpDate" = $9
WHERE "ExpID" = $10;
