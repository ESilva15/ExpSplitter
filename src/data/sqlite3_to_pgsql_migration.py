#!/bin/env python

import psycopg2
from psycopg2 import Error


class DB():
    def __init__(self, user, passwd,
                 db, host, port):
        try:
            self.conn = psycopg2.connect(user=user, password=passwd, database=db,
                                         host=host, port=port)
            self.cursor = self.conn.cursor()
        except (Exception, Error) as error:
            raise error


    def runQuery(self, query, data):
        try:
            print(f"{query} {data}")
            self.cursor.execute(query, data)
            self.conn.commit()
        except (Exception, Error) as error:
            raise error


def get_data(filePath):
    rows = []
    with open(filePath, "r") as f:
        for line in f:
            fields = [v.strip('"') for v in line.strip().split(",")]
            rows.append(fields)

    return rows


def truncate_table(db, m):
    print("Truncating table:", m["table"])
    try:
        db.runQuery(f"TRUNCATE TABLE {m["table"]} CASCADE", [])
    except (Exception, Error) as error:
        raise error


def insert_data(db, query, data):
    for row in data:
        try:
            db.runQuery(query, row)
        except (Exception, Error) as error:
            raise error


def mig_data(db, m):
    print("Running migration for table:", m["table"])

    values = ','.join(['%s'] * len(m["headers"]))
    fields = ','.join(m["headers"])
    query = f"INSERT INTO {m["table"]} ({fields}) VALUES({values})"

    try:
        data = get_data(m["csvFile"])
        insert_data(db, query, data)
    except (Exception, Error) as error:
        raise error


if __name__ == "__main__":
    try:
        dbObj = DB("expuser", "exppass", "expdb", "127.0.0.1", "5431")
    except (Exception, Error) as error:
        print("Threw error:", error)
        exit(0)

    migs = [
        {
            "table": "schema_migrations",
            "csvFile": "./schema_migrations.csv",
            "headers": ["version", "dirty"],
        },
        {
            "table": "categories",
            "csvFile": "./categories.csv",
            "headers": ["CategoryID", "CategoryName"],
        },
        {
            "table": "users",
            "csvFile": "./users.csv",
            "headers": ["UserID", "UserName", "UserPass"],
        },
        {
            "table": "stores",
            "csvFile": "./stores.csv",
            "headers": ["StoreID", "StoreName"],
        },
        {
            "table": "expenseTypes",
            "csvFile": "./expenseTypes.csv",
            "headers": ["TypeID", "TypeName"],
        },
        {
            "table": "expenses",
            "csvFile": "./expenses.csv",
            "headers": ["ExpID", "Description", "Value", "StoreID",
                        "CategoryID", "OwnerUserID", "TypeID", "ExpDate", 
                        "CreationDate", "PaidOff", "SharesEven"],
        },
        {
            "table": "expensesShares",
            "csvFile": "./expensesShares.csv",
            "headers": ["ExpShareID", "ExpID", "UserID", "Share", "Calculated"],
        },
        {
            "table": "expensesPayments",
            "csvFile": "./expensesPayments.csv",
            "headers": ["ExpPaymID", "ExpID", "UserID", "Payed"],
        },
    ]

    for m in migs:
        truncate_table(dbObj, m) 

    for m in migs:
        try:
            mig_data(dbObj, m)
        except (Exception, Error) as error:
            print("failed to migrate data:", error)
