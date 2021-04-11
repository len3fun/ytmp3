.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	source env.sh
	./.bin/bot

build-image:
	docker build -t bot .

start-container:
	docker run --name bot --env-file .env bot