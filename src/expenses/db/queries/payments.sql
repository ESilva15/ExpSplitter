-- name: GetPayments :many
SELECT 
  sqlc.embed(payments), sqlc.embed(users)
FROM 
  "expensesPayments" as payments
JOIN 
  users as users ON users.UserID = payments.UserID
WHERE "ExpID" = $1;

-- name: GetExpensePaymentByUser :one
SELECT
  sqlc.embed(payments), sqlc.embed(users)
FROM
  "expensesPayments" as payments
JOIN
  users as users ON users.UserID = payments.UserID
WHERE
  "ExpID" = $1 AND users."UserID" = $2;
  

-- name: InsertPayment :execresult
INSERT INTO "expensesPayments"(
  "ExpID", "UserID", "Payed"
)
VALUES($1, $2, $3);

-- name: UpdatePayment :execresult
UPDATE "expensesPayments"
SET "UserID" = $1, "Payed" = $2
WHERE "ExpPaymID" = $3;

-- name: DeletePayment :execresult
DELETE FROM "expensesPayments" WHERE "ExpPaymID" = $1;
