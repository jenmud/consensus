// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jenmud/consensus/ent/epic"
	"github.com/jenmud/consensus/ent/predicate"
	"github.com/jenmud/consensus/ent/project"
)

// EpicUpdate is the builder for updating Epic entities.
type EpicUpdate struct {
	config
	hooks    []Hook
	mutation *EpicMutation
}

// Where appends a list predicates to the EpicUpdate builder.
func (eu *EpicUpdate) Where(ps ...predicate.Epic) *EpicUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EpicUpdate) SetName(s string) *EpicUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (eu *EpicUpdate) SetProjectID(id int) *EpicUpdate {
	eu.mutation.SetProjectID(id)
	return eu
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (eu *EpicUpdate) SetNillableProjectID(id *int) *EpicUpdate {
	if id != nil {
		eu = eu.SetProjectID(*id)
	}
	return eu
}

// SetProject sets the "project" edge to the Project entity.
func (eu *EpicUpdate) SetProject(p *Project) *EpicUpdate {
	return eu.SetProjectID(p.ID)
}

// Mutation returns the EpicMutation object of the builder.
func (eu *EpicUpdate) Mutation() *EpicMutation {
	return eu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (eu *EpicUpdate) ClearProject() *EpicUpdate {
	eu.mutation.ClearProject()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EpicUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, EpicMutation](ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EpicUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EpicUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EpicUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EpicUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := epic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Epic.name": %w`, err)}
		}
	}
	return nil
}

func (eu *EpicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   epic.Table,
			Columns: epic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: epic.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(epic.FieldName, field.TypeString, value)
	}
	if eu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   epic.ProjectTable,
			Columns: []string{epic.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   epic.ProjectTable,
			Columns: []string{epic.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{epic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EpicUpdateOne is the builder for updating a single Epic entity.
type EpicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EpicMutation
}

// SetName sets the "name" field.
func (euo *EpicUpdateOne) SetName(s string) *EpicUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (euo *EpicUpdateOne) SetProjectID(id int) *EpicUpdateOne {
	euo.mutation.SetProjectID(id)
	return euo
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (euo *EpicUpdateOne) SetNillableProjectID(id *int) *EpicUpdateOne {
	if id != nil {
		euo = euo.SetProjectID(*id)
	}
	return euo
}

// SetProject sets the "project" edge to the Project entity.
func (euo *EpicUpdateOne) SetProject(p *Project) *EpicUpdateOne {
	return euo.SetProjectID(p.ID)
}

// Mutation returns the EpicMutation object of the builder.
func (euo *EpicUpdateOne) Mutation() *EpicMutation {
	return euo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (euo *EpicUpdateOne) ClearProject() *EpicUpdateOne {
	euo.mutation.ClearProject()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EpicUpdateOne) Select(field string, fields ...string) *EpicUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Epic entity.
func (euo *EpicUpdateOne) Save(ctx context.Context) (*Epic, error) {
	return withHooks[*Epic, EpicMutation](ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EpicUpdateOne) SaveX(ctx context.Context) *Epic {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EpicUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EpicUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EpicUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := epic.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Epic.name": %w`, err)}
		}
	}
	return nil
}

func (euo *EpicUpdateOne) sqlSave(ctx context.Context) (_node *Epic, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   epic.Table,
			Columns: epic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: epic.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Epic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, epic.FieldID)
		for _, f := range fields {
			if !epic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != epic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(epic.FieldName, field.TypeString, value)
	}
	if euo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   epic.ProjectTable,
			Columns: []string{epic.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   epic.ProjectTable,
			Columns: []string{epic.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Epic{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{epic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}