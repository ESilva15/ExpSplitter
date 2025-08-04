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
  "expenseTypes" as types ON types.TypeID = expenses.TypeID;

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
