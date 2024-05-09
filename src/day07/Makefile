.PHONY: test_ex00
test_ex00:
	go test -v ./ex00

.PHONY: test_ex01
test_ex01:
	go test -bench=. -cpuprofile=cpu.prof ./ex01
	go tool pprof cpu.prof

.PHONY: clear
clear:
	@rm -rf cpu.prof ex01.test
