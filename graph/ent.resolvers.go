package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/jenmud/consensus/ent"
	"github.com/jenmud/consensus/graph/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.CommentWhereInput) (*ent.CommentConnection, error) {
	return r.client.Comment.Query().Paginate(ctx, after, first, before, last, ent.WithCommentFilter(where.Filter))
}

// Epics is the resolver for the epics field.
func (r *queryResolver) Epics(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.EpicWhereInput) (*ent.EpicConnection, error) {
	return r.client.Epic.Query().Paginate(ctx, after, first, before, last, ent.WithEpicFilter(where.Filter))
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ProjectWhereInput) (*ent.ProjectConnection, error) {
	return r.client.Project.Query().Paginate(ctx, after, first, before, last, ent.WithProjectFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
