-- name: GetPayments :many
SELECT 
  sqlc.embed(payments), sqlc.embed(users)
FROM 
  "expensesPayments" as payments
JOIN 
  users as users ON users.UserID = payments.UserID
WHERE "ExpID" = ?;

-- name: GetExpensePaymentByUser :one
SELECT
  sqlc.embed(payments), sqlc.embed(users)
FROM
  "expensesPayments" as payments
JOIN
  users as users ON users.UserID = payments.UserID
WHERE
  "ExpID" = ? AND users."UserID" = ?;
  

-- name: InsertPayment :execresult
INSERT INTO "expensesPayments"(
  "ExpID", "UserID", "Payed"
)
VALUES(?, ?, ?);

-- name: UpdatePayment :execresult
UPDATE expensesPayments
SET "UserID" = ?, "Payed" = ?
WHERE "ExpPaymID" = ?;

-- name: DeletePayment :execresult
DELETE FROM "expensesPayments" WHERE "ExpPaymID" = ?;
