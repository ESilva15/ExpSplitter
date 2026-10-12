# Download a production database's data
```bash
docker exec -i <container-name> \
    pg_dump -U <db-user> -a -d <database> > outputfile.sql
```

# Start a local development database


# Upload a database
```bash
# delete the pgsql-data-dev directory contents
# restart the container 
psql -h 127.0.0.1 -p 5431 -U expdbuser -d expenses -f ./data.sql
```
