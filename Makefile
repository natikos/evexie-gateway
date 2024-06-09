build:
	@go build -o ./bin/gateway ./cmd/main.go

run: build
	@./bin/gateway