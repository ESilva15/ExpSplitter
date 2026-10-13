# Download a production database's data
```bash
docker exec -i <container-name> \
    pg_dump -U <db-user> -a -d <database> > outputfile.sql
```

# Start a local development database
```bash
# Use to clean the volumens and kill the current running database
podman compose down -v database
# Use to recreate database
podman compose up database --force-recreate

# Alternatively, to down the service and clear it
make devdb_clear

# And to start it
make devdb_start
```

# Upload a database
```bash
# Don't forget to devdb_clear the current existing database
psql -h 127.0.0.1 -p 5431 -U expdbuser -d expenses -f ./data.sql
```
