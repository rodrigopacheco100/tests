build:
	make docs
	go build -o bin/app main.go

run:
	go run main.go

dev:
	make docs
	go install github.com/mitranim/gow@latest
	gow run ./main.go

docs:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init

debug:
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --continue --output bin/__debug_bin main.go
