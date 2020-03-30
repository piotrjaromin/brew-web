
start:
	go run main.go -type=mock

web:
	npm --prefix web-ui run start

all: start web
