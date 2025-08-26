-- name: GetShares :many
SELECT 
  sqlc.embed(shares), sqlc.embed(users)
FROM 
  "expensesShares" as shares
JOIN 
  users as users ON users.UserID = shares.UserID
WHERE "ExpID" = ?;

-- name: InsertShare :execresult
INSERT INTO "expensesShares"(
  "ExpID", "UserID", "Share", "Calculated"
)
VALUES(?, ?, ?, ?);

-- name: UpdateShare :execresult
UPDATE expensesShares
SET "UserID" = ?, "Share" = ?, "Calculated" = ?
WHERE "ExpShareID" = ?;

-- name: DeleteShare :execresult
DELETE FROM "expensesShares" where "ExpShareID" = ?;
