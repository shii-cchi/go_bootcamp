.PHONY: all
all: build

.PHONY: build
build:
	go build -o server cmd/candy-server-server/main.go

.PHONY: run_server
run_server:
	./server --port=3333

.PHONY: clean
clean:
	@rm -rf server