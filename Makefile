build:
	@go build -o bin/Ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/Ecom