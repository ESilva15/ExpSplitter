-- 0002_change_floats_to_decimal.up.sql

-- expensesShares
CREATE TABLE expensesShares_new (
  "ExpShareID" integer PRIMARY KEY AUTOINCREMENT,
  "ExpID" int NOT NULL,
  "UserID" int NOT NULL,
  "Share" DECIMAL(10,2) NOT NULL DEFAULT 0.5,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID) ON DELETE RESTRICT,
  FOREIGN KEY(UserID) REFERENCES users(UserID) ON DELETE RESTRICT
);
INSERT INTO expensesShares_new SELECT * FROM expensesShares;
DROP TABLE expensesShares;
ALTER TABLE expensesShares_new RENAME TO expensesShares;

-- expensesPayments
CREATE TABLE expensesPayments_new (
  "ExpPaymID" integer PRIMARY KEY AUTOINCREMENT,
  "ExpID" int NOT NULL,
  "UserID" int NOT NULL,
  "Payed" DECIMAL(10,2) NOT NULL DEFAULT 0,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID) ON DELETE RESTRICT,
  FOREIGN KEY(UserID) REFERENCES users(UserID) ON DELETE RESTRICT
);
INSERT INTO expensesPayments_new SELECT * FROM expensesPayments;
DROP TABLE expensesPayments;
ALTER TABLE expensesPayments_new RENAME TO expensesPayments;

-- expenses
CREATE TABLE expenses_new (
  "ExpID" integer PRIMARY KEY AUTOINCREMENT,
  "Description" text NOT NULL,
  "Value" DECIMAL(10,2) NOT NULL DEFAULT 0,
  "StoreID" int NOT NULL DEFAULT 0,
  "CategoryID" int NOT NULL DEFAULT 0,
  "OwnerUserID" int NOT NULL DEFAULT 0,
  "TypeID" int NOT NULL DEFAULT 0,
  "ExpDate" int NOT NULL DEFAULT -1,
  "CreationDate" int NOT NULL DEFAULT -1,
  FOREIGN KEY(StoreID) REFERENCES stores(StoreID) ON DELETE RESTRICT,
  FOREIGN KEY(CategoryID) REFERENCES categories(CategoryID) ON DELETE RESTRICT,
  FOREIGN KEY(OwnerUserID) REFERENCES users(UserID) ON DELETE RESTRICT,
  FOREIGN KEY(TypeID) REFERENCES expensetypes(TypeID) ON DELETE RESTRICT
);
INSERT INTO expenses_new SELECT * FROM expenses;
DROP TABLE expenses;
ALTER TABLE expenses_new RENAME TO expenses;
