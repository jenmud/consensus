all: generate

install-tools:
	go install github.com/air-verse/air@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/spf13/cobra-cli@latest

update-deps:
	go mod tidy

vendor: update-deps
	go mod vendor

generate:
	sqlc generate -f zarf/data/sqlite/sqlc.yml

run:
	mkdir -p ./tmp
	air -c ./zarf/air.toml