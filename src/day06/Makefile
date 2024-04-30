include .env

.PHONY: all
all: create_logo

.PHONY: create_logo
create_logo:
	go build -o create_logo ex00/amazing_logo.go

.PHONY: migration
migration:
	cd ./internal/database/sql/schema && goose postgres ${DB_URI} up

.PHONY: migration_down
migration_down:
	cd ./internal/database/sql/schema && goose postgres ${DB_URI} down

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: clean
clean:
	@rm -rf create_logo