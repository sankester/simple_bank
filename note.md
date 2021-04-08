### run postgres in docker

```
docker run --name postgres12 -p 5433:5433 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

### access psql command

```
docker exec -it postgres12 psql -U root

```

## GO Migrate

### go migrate installation

[Go Migrate Installation](https://stackoverflow.com/questions/61634590/how-to-install-go-migrate-cli-in-linux-or-oracle-linux)

### postgres go migrate tutorial

[Postgres Go Migrate Tutorial](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

### go migarete command
- create migration 
    ```
    migrate create -ext sql -dir db/migrations -seq create_users_table
    ```

## Database

### create new DB

```
createdb --username=root --owner=root simple_bank
```

## SQLC

### Official Repo
[Official Repository](https://github.com/kyleconroy/sqlc)

### installation

[Installation SQLC](https://docs.sqlc.dev/en/latest/overview/install.html)

### command
- sqlc init : init config
- sqlc generate : generating
