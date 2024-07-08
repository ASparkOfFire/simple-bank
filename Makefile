.PHONY: migrate up down reset

up:
	@go run db/migration/cmd/migrate.go up

down:
	@go run db/migration/cmd/migrate.go down

reset:
	@go run db/migration/cmd/migrate.go reset

migration:
	@migrate create -ext sql -seq -dir db/migration $(filter-out $@,$(MAKECMDGOALS))

sqlc:
	@sqlc generate

test:
	@go test -v ./... --count=1 -cover