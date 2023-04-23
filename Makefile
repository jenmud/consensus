all: generate

setup:
	sudo apt update && sudo apt install -y sqlite3 gcc make

generate:
	go generate ./...

run:
	CGO_ENABLED=1 go run . --debug --server 0.0.0.0:8083