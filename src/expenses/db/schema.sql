-- TODO
-- Check why sqlc doesn't correctly parse all my migrations
-- This should only be a workaround :c

CREATE TABLE categories(
  "CategoryID" integer PRIMARY KEY AUTOINCREMENT,
  "CategoryName" text NOT NULL
) 

CREATE TABLE "expenseTypes"(
  "TypeID" integer PRIMARY KEY AUTOINCREMENT,
  "TypeName" text NOT NULL
)

CREATE TABLE "expenses" (
  "ExpID" INTEGER PRIMARY KEY AUTOINCREMENT,
  "Description" TEXT NOT NULL,
  "Value" TEXT NOT NULL DEFAULT 0,
  "StoreID" int NOT NULL DEFAULT 0,
  "CategoryID" int NOT NULL DEFAULT 0,
  "OwnerUserID" int NOT NULL DEFAULT 0,
  "TypeID" int NOT NULL DEFAULT 0,
  "ExpDate" int NOT NULL DEFAULT -1,
  "CreationDate" int NOT NULL DEFAULT -1,
  "PaidOff" BOOL NOT NULL DEFAULT FALSE,
  "SharesEven" BOOL NOT NULL DEFAULT FALSE,
  FOREIGN KEY(StoreID) REFERENCES stores(StoreID) ON DELETE RESTRICT,
  FOREIGN KEY(CategoryID) REFERENCES categories(CategoryID) ON DELETE RESTRICT,
  FOREIGN KEY(OwnerUserID) REFERENCES users(UserID) ON DELETE RESTRICT,
  FOREIGN KEY(TypeID) REFERENCES expensetypes(TypeID) ON DELETE RESTRICT
)

CREATE TABLE "expensesPayments" (
  "ExpPaymID" integer PRIMARY KEY AUTOINCREMENT,
  "ExpID" int NOT NULL,
  "UserID" int NOT NULL,
  "Payed" TEXT NOT NULL DEFAULT 0,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID) ON DELETE RESTRICT,
  FOREIGN KEY(UserID) REFERENCES users(UserID) ON DELETE RESTRICT
)

CREATE TABLE "expensesShares" (
  "ExpShareID" integer PRIMARY KEY AUTOINCREMENT,
  "ExpID" int NOT NULL,
  "UserID" int NOT NULL,
  "Share" TEXT NOT NULL DEFAULT 0.5,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID) ON DELETE RESTRICT,
  FOREIGN KEY(UserID) REFERENCES users(UserID) ON DELETE RESTRICT
)

CREATE TABLE stores(
  "StoreID" integer PRIMARY KEY AUTOINCREMENT,
  "StoreName" text NOT NULL
) 

CREATE TABLE users(
  "UserID" integer PRIMARY KEY AUTOINCREMENT,
  "UserName" text NOT NULL,
  "UserPass" text NOT NULL
)
