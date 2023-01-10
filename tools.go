//go:build tools
// +build tools

package main

import (
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen"
	_ "github.com/spf13/cobra/cobra"
)
