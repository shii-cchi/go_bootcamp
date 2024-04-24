.PHONY: all
all: build_server

.PHONY: build_server
build_server:
	go build -o server cmd/server/main.go

.PHONY: clean
clean:
	@rm -rf server

.PHONY: protoc_gen
protoc_gen:
	protoc -I internal/proto/ --go_out=internal --go-grpc_out=internal internal/proto/transmitter.proto