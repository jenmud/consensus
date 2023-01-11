package graph

import "github.com/jenmud/consensus/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

// New returns a new graph resolver.
func New(client *ent.Client) *Resolver {
	return &Resolver{client: client}
}
