build:
	@go build -o bin/yunque

run: build
	@./bin/yunque

test:
	@go test -v ./...