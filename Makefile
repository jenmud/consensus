all: generate

generate:
	go generate ./...

run:
	go run . --debug --address localhost:8003