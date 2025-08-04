-- name: GetPayments :many
SELECT 
  sqlc.embed(payments), sqlc.embed(users)
FROM 
  "expensesPayments" as payments
JOIN 
  users as users ON users.UserID = payments.UserID
WHERE "ExpID" = ?;
