package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin implements the ent.Mixin for sharing time fields with package schemas.
type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			).
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			).
			UpdateDefault(time.Now),
	}
}

// TaskMixin implements the ent.Mixin for sharing task type fields with package schemas.
type TaskMixin struct {
	mixin.Schema
}

func (TaskMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deadline").
			Annotations(
				entgql.OrderField("DEADLINE"),
			),

		field.Enum("size").
			NamedValues("small", "SMALL", "medium", "MEDIUM", "large", "LARGE").
			Default("small").
			Annotations(
				entgql.OrderField("SIZE"),
			),

		field.Enum("status").
			NamedValues("pending", "PENDING", "in_progress", "IN_PROGRESS", "done", "DONE").
			Default("pending").
			Annotations(
				entgql.OrderField("STATUS"),
			),
	}
}
