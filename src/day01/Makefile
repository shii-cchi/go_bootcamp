.PHONY: all
all: build_ex00 build_ex01 build_ex02

.PHONY: build_ex00
build_ex00:
	go build -o readDB cmd/dbreader/main.go

.PHONY: build_ex01
build_ex01:
	go build -o compareDB cmd/dbcomparer/main.go

.PHONY: build_ex02
build_ex02:
	go build -o compareFS cmd/fscomparer/main.go

.PHONY: clean
clean:
	@rm -rf readDB compareDB compareFS