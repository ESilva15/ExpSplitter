-- 0002_change_floats_to_decimal.down.sql

-- expenses
CREATE TABLE expenses_old (
  "ExpID" integer PRIMARY KEY AUTOINCREMENT,
  "Description" text NOT NULL,
  "Value" float NOT NULL DEFAULT 0,
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
INSERT INTO expenses_old SELECT * FROM expenses;
DROP TABLE expenses;
ALTER TABLE expenses_old RENAME TO expenses;

-- expensesShares
CREATE TABLE expensesShares_old (
  "ExpShareID" integer PRIMARY KEY AUTOINCREMENT,
  "ExpID" int NOT NULL,
  "UserID" int NOT NULL,
  "Share" float NOT NULL DEFAULT 0.5,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID) ON DELETE RESTRICT,
  FOREIGN KEY(UserID) REFERENCES users(UserID) ON DELETE RESTRICT
);
INSERT INTO expensesShares_old SELECT * FROM expensesShares;
DROP TABLE expensesShares;
ALTER TABLE expensesShares_old RENAME TO expensesShares;

-- expensesPayments
CREATE TABLE expensesPayments_old (
  "ExpPaymID" integer PRIMARY KEY AUTOINCREMENT,
  "ExpID" int NOT NULL,
  "UserID" int NOT NULL,
  "Payed" float NOT NULL DEFAULT 0,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID) ON DELETE RESTRICT,
  FOREIGN KEY(UserID) REFERENCES users(UserID) ON DELETE RESTRICT
);
INSERT INTO expensesPayments_old SELECT * FROM expensesPayments;
DROP TABLE expensesPayments;
ALTER TABLE expensesPayments_old RENAME TO expensesPayments;
