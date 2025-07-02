build:
	@go build -o bin/mugi-id-server main.go

test:
	@go test -v ./...

run: build
	@./bin/mugi-id-server
	