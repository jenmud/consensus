package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Epic holds the schema definition for the Epic entity.
type Epic struct {
	ent.Schema
}

// Annotations of the Epic.
func (Epic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Mixins of the Epic.
func (Epic) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		TaskMixin{},
	}
}

// Fields of the Epic.
func (Epic) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.String("description"),
	}
}

// Edges of the Epic.
func (Epic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("project", Project.Type).Unique(),
		edge.From("reporter", User.Type).Ref("reporter").Unique(),
		edge.From("assignee", User.Type).Ref("assignee").Unique(),
		edge.To("comments", Comment.Type),
	}
}
