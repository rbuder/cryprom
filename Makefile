# Usage:
# make        # compile all binary
# make clean  # remove ALL binaries and objects
.DEFAULT_GOAL := build

build:
	@echo "Building Linux binary"
	GOOS=linux GOARCH=amd64 go build -o bin/cryprom

clean:
	@echo "Cleaning up"
	rm -rf ./bin

image:
	docker build .