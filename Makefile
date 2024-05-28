IP_COUNT ?= 3
URL_COUNT ?= 3

.PHONY: all
all: tidy build test

.PHONY: clean
clean:
	rm -rf app/*

.PHONY: build
build:
	go build -o ./app/logparser ./cmd/logparser/main.go

.PHONY: run-all 
run-all: 
	./app/logparser log -- url top $(URL_COUNT) - ip active $(IP_COUNT) - ip unique

.PHONY: run-ip
run-ip: 
	./app/logparser log -- ip active $(IP_COUNT) - ip unique

.PHONY: run-ip-active 
run-ip-active: 
	@echo "Usage: make run-ip IP_COUNT=<number>"
	./app/logparser log -- ip active $(IP_COUNT)

.PHONY: run-ip-unique 
run-ip-unique: 
	./app/logparser log -- ip unique

.PHONY: run-url-top
run-url-top: 
	@echo "Usage: make run-url URL_COUNT=<number>"
	./app/logparser log -- url top $(URL_COUNT)

.PHONY: run-counts
run-counts: 
	@echo "Usage: make run-url URL_COUNT=<number>"
	./app/logparser log -- ip active $(IP_COUNT) - url top $(URL_COUNT)

.PHONY: run-help
run-help: 
	@echo "Usage: make run-all IP_COUNT=x URL_COUNT=x \n"
	./app/logparser log --help

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -v ./internal/...

.PHONY: bench-all
bench-all: 
	go test -bench=. ./...
