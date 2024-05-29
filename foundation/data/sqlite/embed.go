package sqlite

import "embed"

//go:embed schema.sql
var schema embed.FS
