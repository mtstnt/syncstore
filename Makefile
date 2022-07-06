.PHONY: build

build:
	cd app && \
	go build -o ../build/ .

run: build
	./build/syncstore

test:
	cd app && \
	go test ./...

wipe-db: 
	rm -f database/my.db && \
	touch database/my.db