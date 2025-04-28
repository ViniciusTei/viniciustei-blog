run:
	go run ./cmd/blog

build:
	go build -o dist/server ./cmd/blog

test:
	go test -v ./...

coverage:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
