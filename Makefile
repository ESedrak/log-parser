.PHONY: all
all: tidy build test

.PHONY: clean
clean:
	rm -rf app/*

.PHONY: build
build:
	go build -o ./app/logparser ./cmd/logparser/main.go

.PHONY: run
run: 
	./app/logparser

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -v ./internal/...

.PHONY: bench-all
bench-all: 
	go test -bench=. ./...
