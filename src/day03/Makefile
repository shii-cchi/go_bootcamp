.PHONY: all
all: load_data search_app

.PHONY: load_data
load_data:
	go build -o load_data cmd/load_data/main.go

.PHONY: search_app
search_app:
	go build -o search_app cmd/search_app/main.go

.PHONY: clean
clean:
	@rm -rf load_data search_app