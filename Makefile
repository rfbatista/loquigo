test:
	go test ./... -v
run:
	go run main.go
wire:
	wire ./cmd
build:
	docker-compose up --build
dev:
	docker-compose -f docker-compose.yaml up --build
debug:
	docker-compose -f docker-compose.debug.yaml up --build