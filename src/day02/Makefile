.PHONY: all
all: build_ex00 build_ex01 build_ex02 build_ex03

.PHONY: build_ex00
build_ex00:
	go build -o myFind cmd/finder/main.go

.PHONY: build_ex01
build_ex01:
	go build -o myWc cmd/counter/main.go

.PHONY: build_ex02
build_ex02:
	go build -o myXargs cmd/runner/main.go

.PHONY: build_ex03
build_ex03:
	go build -o myRotate cmd/fscomparer/main.go

.PHONY: clean
clean:
	@rm -rf myFind myWc myXargs myRotate