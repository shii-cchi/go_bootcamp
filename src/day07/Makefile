.PHONY: test_ex00
test_ex00:
	go test -v ./ex00

.PHONY: test_ex01
test_ex01:
	go test -bench=. -cpuprofile=cpu.prof ./ex01
	go tool pprof cpu.prof

.PHONY: doc
doc:
	godoc -http=:6060

.PHONY: save_doc
save_doc:
	godoc -url http://localhost:6060/pkg/day07/ex00 > doc.html
	zip docs.zip doc.html

.PHONY: clear
clear:
	@rm -rf cpu.prof ex01.test doc.html docs.zip
