#!/bin/sh
set -e

if [ -z "$(ls -A "$PGDATA")" ]; then
  echo "Creating database with:"
  echo "DIR: $PGDATA"
  echo "USER: $ExpDB_USER"
  echo "PASS: $ExpDB_PASS"
  echo "NAME: $ExpDB_NAME"
  # Initiate PG data directory
  initdb \
    --locale=C.UTF-8 \
    --encoding=UTF8 \
    -D "$PGDATA" \
    --data-checksums \
    --username="$UNAME"

  # Set PG to listen on all addresses
  sed -ri "s/^#(listen_addresses\s*=\s*)\S+/\1'*'/" "$PGDATA"/postgresql.conf

  # Start PG to add the NC DB
  pg_ctl -D "$PGDATA" \
    -o "-c listen_addresses=''" \
    -w start

  # Create the NC database
  createUser="CREATE USER $ExpDB_USER WITH PASSWORD '$ExpDB_PASS';"
  createDB="CREATE DATABASE $ExpDB_NAME TEMPLATE template0 ENCODING 'UNICODE';"
  alterDB="ALTER DATABASE $ExpDB_NAME OWNER TO $ExpDB_USER;"
  grantPriv="GRANT ALL PRIVILEGES ON DATABASE $ExpDB_NAME TO $ExpDB_USER;"
  psql -d postgres -c "${createUser}"
  psql -d postgres -c "${createDB}"
  psql -d postgres -c "${alterDB}"
  psql -d postgres -c "${grantPriv}"

  # Execute the initial schema
  psql -U "$ExpDB_USER" -d "$ExpDB_NAME" -a -f /schema.sql

  pg_ctl -D "$PGDATA" -m fast -w stop

  { echo; echo "host all all 0.0.0.0/0 md5"; } >> "$PGDATA"/pg_hba.conf
fi

postgres "$@"
