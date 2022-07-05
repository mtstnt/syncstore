build:
	go build -o build/ .

run: build
	./build/syncstore

test:
	go test ./...