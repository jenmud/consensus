PROTOC_VERSION=27.1
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip
PROTOC_PATH=~/.local/concencus/protoc
GOROOT := $(shell go env GOROOT)
GOPATH := $(shell go env GOPATH)
PATH=$(GOROOT)/bin:$(GOPATH)/bin:$(PROTOC_PATH)/bin:/usr/bin:/usr/local/bin:$$PATH
TEMP := $(shell /usr/bin/mktemp -d -u)

all: generate

install: download-and-install-protoc install-tools

download-and-install-protoc:
	PATH=$(PATH) /usr/bin/mkdir -p $(PROTOC_PATH)
	PATH=$(PATH) /usr/bin/curl -L $(PROTOC_URL) -o $(TEMP)/protoc.zip
	PATH=$(PATH) /usr/bin/unzip -u $(TEMP)/protoc.zip -d $(PROTOC_PATH)
	PATH=$(PATH) /usr/bin/rm -rf $(TEMP)

install-tools:
	PATH=$(PATH) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	PATH=$(PATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	PATH=$(PATH) go install github.com/air-verse/air@latest
	PATH=$(PATH) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	PATH=$(PATH) go install github.com/spf13/cobra-cli@latest

update-deps:
	PATH=$(PATH) go mod tidy

vendor: update-deps
	PATH=$(PATH) go mod vendor

generate-sql:
	PATH=$(PATH) sqlc generate -f zarf/data/sqlite/sqlc.yml

generate-proto:
	PATH=$(PATH) /usr/bin/mkdir -p ./business/service
	cd ./zarf/proto && PATH=$(PATH) protoc \
	--go_out=../../business/service \
	--go_opt=paths=source_relative \
    --go-grpc_out=../../business/service \
	--go-grpc_opt=paths=source_relative \
   service.proto

generate: generate-proto generate-sql

run-ui:
	PATH=$(PATH) /usr/bin/mkdir -p ./tmp
	PATH=$(PATH) air -c ./zarf/air.toml

run-service:
	# needs cgo enabled because of sqlite3
	PATH=$(PATH) CGO_ENABLED=1 go run ./app