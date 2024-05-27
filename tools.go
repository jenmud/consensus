//go:build tools
// +build tools

package main

import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
