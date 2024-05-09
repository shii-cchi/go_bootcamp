include blog/.env

.PHONY: all
all: create_logo build_server

.PHONY: create_logo
create_logo:
	go build -o create_logo logo/amazing_logo.go

.PHONY: build_server
build_server:
	go build -o server blog/cmd/main.go

.PHONY: migration
migration:
	cd blog/internal/database/schema && goose postgres ${DB_URI} up

.PHONY: migration_down
migration_down:
	cd blog/internal/database/schema && goose postgres ${DB_URI} down

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test_rate_limiter
test_rate_limiter:
	for i in $$(seq 1 200); do curl -X GET http://localhost:8888/1 & sleep 0.0001; done

.PHONY: clean
clean:
	@rm -rf create_logo server