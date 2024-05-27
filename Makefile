all: generate

vendor:
	go mod vendor

run:
	go run . --debug --server 0.0.0.0:8083