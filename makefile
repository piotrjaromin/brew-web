
install:
	go get github.com/rakyll/statik
	go mod tidy
	go mod download


test:
	go test ./...
run:
	go run main.go -type=mock

web:
	npm --prefix web-ui run start

all: start web
