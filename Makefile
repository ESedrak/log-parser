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

.PHONY: run-all 
run-all: 
	./app/logparser log -- ip unique - ip active $(NUM_IPS) - url top $(NUM_URLS)

.PHONY: run-ip
run-ip: 
	./app/logparser log -- ip unique - ip active $(NUM_IPS)

.PHONY: run-ip-active 
run-ip-active: 
	@echo "Usage: make run-ip NUM_IPS=<number>"
	./app/logparser log -- ip active $(NUM_IPS)

.PHONY: run-ip-unique 
run-ip-unique: 
	./app/logparser log -- ip unique

.PHONY: run-url-top
run-url-top: 
	@echo "Usage: make run-url NUM_URLS=<number>"
	./app/logparser log -- url top $(NUM_URLS)

.PHONY: run-counts
run-counts: 
	@echo "Usage: make run-url NUM_URLS=<number>"
	./app/logparser log -- ip active $(NUM_IPS) - url top $(NUM_URLS)

.PHONY: run-help
run-help: 
	@echo "Usage: make run-all NUM_IPS=x NUM_URLS=x \n"
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
