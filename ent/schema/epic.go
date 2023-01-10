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
	}
}

// Fields of the Epic.
func (Epic) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Epic.
func (Epic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("project", Project.Type).Unique(),
	}
}
