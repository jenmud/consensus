package sqlite

import "embed"

//go:embed schema.sql
var Schema embed.FS
