.PHONY: build
build:
	go build cmd/main.go

.PHONY: run
run: build
	./main

.PHONY: test
test:
	go test -v tests/main_test.go

.PHONY: test_statistics
test_statistics:
	go test -v tests/statistics_test.go

.PHONY: clean
clean:
	@rm -rf main