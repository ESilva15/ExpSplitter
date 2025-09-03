#!/bin/env python

import psycopg2
from psycopg2 import Error
from datetime import datetime


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


def convert_to_date(timestamp):
    dt = datetime.fromtimestamp(int(timestamp))
    dt = dt.replace(hour=0)
    dt = dt.replace(minute=0)
    dt = dt.replace(second=0)
    return dt.strftime("%Y-%m-%d %H:%M:%S")


def migrate_dates():
    expDateIdx = 7
    creationDateIdx = 8

    newFile = open("./new_expenses.csv", "w")

    with open("./expenses.csv.bk", "r") as f:
        for l in f:
            fields = l.strip().split(',')
            exp = fields[expDateIdx]
            creation = fields[creationDateIdx]
            
            exp = convert_to_date(exp)
            creation = convert_to_date(creation) if creation != -1 else convert_to_date(exp)

            fields[expDateIdx] = exp
            fields[creationDateIdx] = creation
            newFile.write(','.join(fields) + "\n")


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
        db.runQuery(f"TRUNCATE TABLE \"{m["table"]}\" CASCADE", [])
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
    fields = ','.join([f"\"{h}\"" for h in m["headers"]])
    query = f"INSERT INTO \"{m["table"]}\" ({fields}) VALUES({values})"

    lastId = -1
    try:
        data = get_data(m["csvFile"])
        insert_data(db, query, data)
        lastId = data[0]
    except (Exception, Error) as error:
        raise error

    if lastId != -1:
        if m["table"] == "schema_migrations":
            return
        try:
            db.runQuery(
                f"SELECT setval(pg_get_serial_sequence('\"{m["table"]}\"', '{m["headers"][0]}'),"
                f"(SELECT MAX(\"{m["headers"][0]}\") FROM \"{m["table"]}\"))", []
            )
        except (Exception, Error) as error:
            raise error

if __name__ == "__main__":
    if False:
        migrate_dates()

    if True:
        try:
            dbObj = DB("expuser", "exppass", "expdb", "127.0.0.1", "5431")
        except (Exception, Error) as error:
            print("Threw error:", error)
            exit(0)

        migs = [
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
                exit(1)
