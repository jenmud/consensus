package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Annotations of the Project.
func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Mixins of the Project.
func (Project) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("epics", Epic.Type).Ref("project"),
		edge.From("reporter", User.Type).Ref("reporter").Unique(),
		edge.From("assignee", User.Type).Ref("assignee").Unique(),
	}
}
