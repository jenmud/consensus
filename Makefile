all: generate

update-deps:
	go mod tidy

vendor: update-deps
	go mod vendor

generate:
	sqlc generate -f zarf/data/sqlite/sqlc.yml

run:
	go run . --debug --server 0.0.0.0:8083