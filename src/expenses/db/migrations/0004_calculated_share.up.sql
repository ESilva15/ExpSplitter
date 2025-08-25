-- 0004_calculated_share.up.sql

ALTER TABLE "expensesShares" ADD COLUMN "Calculated" TEXT NOT NULL DEFAULT 0;
