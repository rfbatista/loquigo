setup:
	go mod tidy

build:
	cd engine && go build .

run:
	go run .
