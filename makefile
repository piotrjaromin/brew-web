
install:
	go install github.com/rakyll/statik@v0.1.7
	go mod tidy
	go mod download

make build:
	statik -src ./web-ui/build
	go build -o ./bin/brew-web ./main.go

test:
	go test ./...
run:
	go run main.go -type=mock

web:
	npm --prefix web-ui run start

all: start web
