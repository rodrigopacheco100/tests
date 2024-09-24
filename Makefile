build:
	go build -o bin/app src/main.go

run:
	go run src/main.go

dev:
	go install github.com/mitranim/gow@latest
	gow run ./src/main.go
