-- 0003_add_paidOff_column.up.sql

ALTER TABLE expenses ADD COLUMN "PaidOff" BOOL DEFAULT FALSE;
ALTER TABLE expenses ADD COLUMN "SharesEven" BOOL DEFAULT FALSE;

-- Update the SharesEven column
UPDATE expenses
SET SharesEven = TRUE
WHERE ExpID IN (
  SELECT e.ExpID
  FROM expenses e
  JOIN expensesPayments p ON p.ExpID = e.ExpID
  GROUP BY e.ExpID
  HAVING ROUND(SUM(CAST(p.Payed AS REAL)), 2) == e.Value
);

-- Update the PaidOff column
UPDATE expenses
SET "PaidOff" = TRUE
WHERE ExpID IN (
  SELECT p.ExpID as ExpID
  FROM expensesPayments AS p 
  JOIN expenses as e
    ON e.ExpID = p.ExpID
  GROUP BY p.ExpId
  HAVING ROUND(CAST(e.Value AS REAL),2) == ROUND(SUM(CAST(p.Payed AS REAL)), 2)
);
