// Code generated by ent, DO NOT EDIT.

package epic

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/jenmud/consensus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Epic {
	return predicate.Epic(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Epic {
	return predicate.Epic(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Epic {
	return predicate.Epic(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Epic {
	return predicate.Epic(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Epic {
	return predicate.Epic(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Epic {
	return predicate.Epic(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Epic {
	return predicate.Epic(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Epic {
	return predicate.Epic(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Epic {
	return predicate.Epic(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Epic {
	return predicate.Epic(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Epic {
	return predicate.Epic(sql.FieldEQ(FieldDescription, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Epic {
	return predicate.Epic(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Epic {
	return predicate.Epic(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Epic {
	return predicate.Epic(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Epic {
	return predicate.Epic(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Epic {
	return predicate.Epic(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Epic {
	return predicate.Epic(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Epic {
	return predicate.Epic(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Epic {
	return predicate.Epic(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Epic {
	return predicate.Epic(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Epic {
	return predicate.Epic(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Epic {
	return predicate.Epic(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Epic {
	return predicate.Epic(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Epic {
	return predicate.Epic(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Epic {
	return predicate.Epic(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Epic {
	return predicate.Epic(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Epic {
	return predicate.Epic(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Epic {
	return predicate.Epic(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Epic {
	return predicate.Epic(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Epic {
	return predicate.Epic(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Epic {
	return predicate.Epic(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Epic {
	return predicate.Epic(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Epic {
	return predicate.Epic(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Epic {
	return predicate.Epic(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Epic {
	return predicate.Epic(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Epic {
	return predicate.Epic(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Epic {
	return predicate.Epic(sql.FieldContainsFold(FieldDescription, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReporter applies the HasEdge predicate on the "reporter" edge.
func HasReporter() predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReporterTable, ReporterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReporterWith applies the HasEdge predicate on the "reporter" edge with a given conditions (other predicates).
func HasReporterWith(preds ...predicate.User) predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReporterInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReporterTable, ReporterColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssignee applies the HasEdge predicate on the "assignee" edge.
func HasAssignee() predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssigneeTable, AssigneeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssigneeWith applies the HasEdge predicate on the "assignee" edge with a given conditions (other predicates).
func HasAssigneeWith(preds ...predicate.User) predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssigneeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssigneeTable, AssigneeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CommentsTable, CommentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CommentsTable, CommentsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Epic) predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Epic) predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Epic) predicate.Epic {
	return predicate.Epic(func(s *sql.Selector) {
		p(s.Not())
	})
}
