all: generate

generate:
	go generate ./...

run:
	go run . --debug --server 0.0.0.0:8083