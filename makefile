start:
	go run main.go -type=esp

web:
	npm --prefix web-ui run start

all:
	start && web &&