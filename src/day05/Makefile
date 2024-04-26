.PHONY: all
all: build

.PHONY: build
build:
	go build cmd/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	@rm -rf main