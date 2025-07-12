build:
	@go build -o bin/mugi-id-server cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/mugi-id-server
	