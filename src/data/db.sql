DROP TABLE IF EXISTS expenses;
CREATE TABLE expenses (
  ExpID integer PRIMARY KEY AUTOINCREMENT,
  Description text NOT NULL,
  Value float NOT NULL DEFAULT 0,
  StoreID int NOT NULL DEFAULT 0,
  CategoryID int NOT NULL DEFAULT 0,
  OwnerUserID int NOT NULL DEFAULT 0,
  ExpDate int NOT NULL DEFAULT -1,
  CreationDate int NOT NULL DEFAULT -1,
  FOREIGN KEY(StoreID) REFERENCES stores(StoreID),
  FOREIGN KEY(CategoryID) REFERENCES categories(CategoryID),
  FOREIGN KEY(OwnerUserID) REFERENCES users(UserID)
);

DROP TABLE IF EXISTS stores;
CREATE TABLE stores(
  StoreID integer PRIMARY KEY AUTOINCREMENT,
  StoreName text NOT NULL
);
-- Create a default store for NoStore
INSERT INTO stores (StoreID, StoreName) VALUES(0, 'No Store');

DROP TABLE IF EXISTS categories;
CREATE TABLE categories(
  CategoryID integer PRIMARY KEY AUTOINCREMENT,
  CategoryName text NOT NULL
);
-- Create a default category for NoCategory
INSERT INTO categories (CategoryID, CategoryName) VALUES(0, 'No Category');

DROP TABLE IF EXISTS users;
CREATE TABLE users(
  UserID integer PRIMARY KEY AUTOINCREMENT,
  UserName text NOT NULL,
  UserPass text NOT NULL
);

DROP TABLE IF EXISTS expensesShares;
CREATE TABLE expensesShares(
  ExpShareID integer PRIMARY KEY AUTOINCREMENT,
  ExpID int NOT NULL,
  UserID int NOT NULL,
  Share float NOT NULL DEFAULT 0.5,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID),
  FOREIGN KEY(UserID) REFERENCES users(UserID)
);

DROP TABLE IF EXISTS expensesPayments;
CREATE TABLE expensesPayments(
  ExpPaymID integer PRIMARY KEY AUTOINCREMENT,
  ExpID int NOT NULL,
  UserID int NOT NULL,
  Payed float NOT NULL DEFAULT 0,
  FOREIGN KEY(ExpID) REFERENCES expenses(ExpID),
  FOREIGN KEY(UserID) REFERENCES users(UserID)
);
