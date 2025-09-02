-- name: GetShares :many
SELECT 
  sqlc.embed(shares), sqlc.embed(users)
FROM 
  "expensesShares" as shares
JOIN 
  users as users ON users.UserID = shares.UserID
WHERE "ExpID" = $1;

-- name: InsertShare :execresult
INSERT INTO "expensesShares"(
  "ExpID", "UserID", "Share", "Calculated"
)
VALUES($1, $2, $3, $4);

-- name: UpdateShare :execresult
UPDATE "expensesShares"
SET "UserID" = $1, "Share" = $2, "Calculated" = $3
WHERE "ExpShareID" = $4;

-- name: DeleteShare :execresult
DELETE FROM "expensesShares" where "ExpShareID" = $1;
