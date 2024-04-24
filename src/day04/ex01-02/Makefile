.PHONY: all
all: build_server build_client

.PHONY: build_server
build_server:
	go build -o server cmd/candy-server-server/main.go

.PHONY: build_client
build_client:
	go build -o client cmd/test_client/main.go

.PHONY: clean
clean:
	@rm -rf server client