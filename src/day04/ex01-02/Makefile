.PHONY: all
all: build_server build_client

.PHONY: build_server
build_server:
	go build -o server cmd/candy-server-server/main.go

.PHONY: build_client
build_client:
	go build -o client cmd/test_client/main.go

.PHONY: run_server
run_server:
	./server --tls-certificate ca/candy.tld/cert.pem --tls-key ca/candy.tld/key.pem --tls-port=3333

.PHONY: clean
clean:
	@rm -rf server client