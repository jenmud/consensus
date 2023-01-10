all: generate

generate:
	go generate ./...

run:
	go run . --debug 