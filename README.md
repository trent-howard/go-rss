We use `goose` to manage migrations and `sqlc` to compile SQL queries into Go functions because that's what the tute uses. Pretty cool tools, but it'd be nice to have a more DB agnostic way to define these schemas and not have to worry about little syntax differences between, say, Postgres and SQLite. It's great not having to write your own database methods though!

We can install them with `brew` or using `go install`

```sh
brew install goose sqlc
# OR
go install github.com/pressly/goose/v3/cmd/goose@latest github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

To init the database and run any pending migrations navigate to the `./sql/schema` folder and run `goose`

```sh
cd ./sql/schema
goose sqlite3 ../../local.db up
```

If you make any changes to the sql queries you'll need to re-generate the related Go functions. From the project root run

```sh
sqlc generate
```
