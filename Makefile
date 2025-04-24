run:
	go run *.go

build:
	go build -o dist/server

test:
	go test ./...

coverage:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
