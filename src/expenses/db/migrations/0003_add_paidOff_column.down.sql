-- 0003_add_paidOff_column.down.sql

ALTER TABLE expenses DROP COLUMN "PaidOff";
ALTER TABLE expenses DROP COLUMN "SharesEven";
