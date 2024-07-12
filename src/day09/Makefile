.PHONY: all
all: test_ex00 test_ex01 test_ex02

.PHONY: test_ex00
test_ex00:
	go test -v ./ex00

.PHONY: test_ex01
test_ex01:
	go build -o crawlweb ex01/crawlweb.go && ./crawlweb

.PHONY: test_ex02
test_ex02:
	go test -v ./ex02
