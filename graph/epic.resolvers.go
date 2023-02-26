package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/jenmud/consensus/ent"
)

// CreateEpic is the resolver for the createEpic field.
func (r *mutationResolver) CreateEpic(ctx context.Context, input ent.CreateEpicInput) (*ent.Epic, error) {
	return r.client.Epic.Create().SetInput(input).Save(ctx)
}