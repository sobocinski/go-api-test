## Start API
Run: `go run cmd/server/main.go`


## Migrations
Library https://github.com/golang-migrate/migrate    
Install on OS X `brew install golang-migrate`

1. Create new migrations files:
`migrate create -ext sql -dir db/migrations -seq MIGRATION_NAME`

2. Runing migrations (todo env)
`migrate -path db/migrations -database "postgres://user:pass@localhost:5432/db_name?sslmode=disable" -verbose up`    
or     
`migrate -path db/migrations -database "postgres://user:pass@localhost:5432/db_name?sslmode=disable" -verbose down`
