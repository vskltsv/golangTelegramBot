.PHONY:
.SILENT:
build:
	go build -o bot cmd/bot/main.go

run:build
	./bot