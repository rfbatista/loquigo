setup:
	go mod tidy

build:
	cd engine && go build .

run:
	cd engine/cmd && go run .
