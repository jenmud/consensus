// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jenmud/consensus/ent/comment"
	"github.com/jenmud/consensus/ent/epic"
	"github.com/jenmud/consensus/ent/predicate"
	"github.com/jenmud/consensus/ent/project"
	"github.com/jenmud/consensus/ent/user"
)

// CommentQuery is the builder for querying Comment entities.
type CommentQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	inters            []Interceptor
	predicates        []predicate.Comment
	withEpics         *EpicQuery
	withProjects      *ProjectQuery
	withUsers         *UserQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Comment) error
	withNamedEpics    map[string]*EpicQuery
	withNamedProjects map[string]*ProjectQuery
	withNamedUsers    map[string]*UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommentQuery builder.
func (cq *CommentQuery) Where(ps ...predicate.Comment) *CommentQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CommentQuery) Limit(limit int) *CommentQuery {
	cq.limit = &limit
	return cq
}

// Offset to start from.
func (cq *CommentQuery) Offset(offset int) *CommentQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CommentQuery) Unique(unique bool) *CommentQuery {
	cq.unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CommentQuery) Order(o ...OrderFunc) *CommentQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryEpics chains the current query on the "epics" edge.
func (cq *CommentQuery) QueryEpics() *EpicQuery {
	query := (&EpicClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, selector),
			sqlgraph.To(epic.Table, epic.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, comment.EpicsTable, comment.EpicsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProjects chains the current query on the "projects" edge.
func (cq *CommentQuery) QueryProjects() *ProjectQuery {
	query := (&ProjectClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, comment.ProjectsTable, comment.ProjectsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (cq *CommentQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, comment.UsersTable, comment.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Comment entity from the query.
// Returns a *NotFoundError when no Comment was found.
func (cq *CommentQuery) First(ctx context.Context) (*Comment, error) {
	nodes, err := cq.Limit(1).All(newQueryContext(ctx, TypeComment, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{comment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommentQuery) FirstX(ctx context.Context) *Comment {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Comment ID from the query.
// Returns a *NotFoundError when no Comment ID was found.
func (cq *CommentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(newQueryContext(ctx, TypeComment, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{comment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CommentQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Comment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Comment entity is found.
// Returns a *NotFoundError when no Comment entities are found.
func (cq *CommentQuery) Only(ctx context.Context) (*Comment, error) {
	nodes, err := cq.Limit(2).All(newQueryContext(ctx, TypeComment, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{comment.Label}
	default:
		return nil, &NotSingularError{comment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommentQuery) OnlyX(ctx context.Context) *Comment {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Comment ID in the query.
// Returns a *NotSingularError when more than one Comment ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CommentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(newQueryContext(ctx, TypeComment, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{comment.Label}
	default:
		err = &NotSingularError{comment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CommentQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Comments.
func (cq *CommentQuery) All(ctx context.Context) ([]*Comment, error) {
	ctx = newQueryContext(ctx, TypeComment, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Comment, *CommentQuery]()
	return withInterceptors[[]*Comment](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommentQuery) AllX(ctx context.Context) []*Comment {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Comment IDs.
func (cq *CommentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeComment, "IDs")
	if err := cq.Select(comment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CommentQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CommentQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeComment, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CommentQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommentQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeComment, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommentQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommentQuery) Clone() *CommentQuery {
	if cq == nil {
		return nil
	}
	return &CommentQuery{
		config:       cq.config,
		limit:        cq.limit,
		offset:       cq.offset,
		order:        append([]OrderFunc{}, cq.order...),
		inters:       append([]Interceptor{}, cq.inters...),
		predicates:   append([]predicate.Comment{}, cq.predicates...),
		withEpics:    cq.withEpics.Clone(),
		withProjects: cq.withProjects.Clone(),
		withUsers:    cq.withUsers.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithEpics tells the query-builder to eager-load the nodes that are connected to
// the "epics" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CommentQuery) WithEpics(opts ...func(*EpicQuery)) *CommentQuery {
	query := (&EpicClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withEpics = query
	return cq
}

// WithProjects tells the query-builder to eager-load the nodes that are connected to
// the "projects" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CommentQuery) WithProjects(opts ...func(*ProjectQuery)) *CommentQuery {
	query := (&ProjectClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withProjects = query
	return cq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CommentQuery) WithUsers(opts ...func(*UserQuery)) *CommentQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withUsers = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Comment.Query().
//		GroupBy(comment.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CommentQuery) GroupBy(field string, fields ...string) *CommentGroupBy {
	cq.fields = append([]string{field}, fields...)
	grbuild := &CommentGroupBy{build: cq}
	grbuild.flds = &cq.fields
	grbuild.label = comment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Comment.Query().
//		Select(comment.FieldText).
//		Scan(ctx, &v)
func (cq *CommentQuery) Select(fields ...string) *CommentSelect {
	cq.fields = append(cq.fields, fields...)
	sbuild := &CommentSelect{CommentQuery: cq}
	sbuild.label = comment.Label
	sbuild.flds, sbuild.scan = &cq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommentSelect configured with the given aggregations.
func (cq *CommentQuery) Aggregate(fns ...AggregateFunc) *CommentSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CommentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.fields {
		if !comment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Comment, error) {
	var (
		nodes       = []*Comment{}
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withEpics != nil,
			cq.withProjects != nil,
			cq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Comment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Comment{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withEpics; query != nil {
		if err := cq.loadEpics(ctx, query, nodes,
			func(n *Comment) { n.Edges.Epics = []*Epic{} },
			func(n *Comment, e *Epic) { n.Edges.Epics = append(n.Edges.Epics, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withProjects; query != nil {
		if err := cq.loadProjects(ctx, query, nodes,
			func(n *Comment) { n.Edges.Projects = []*Project{} },
			func(n *Comment, e *Project) { n.Edges.Projects = append(n.Edges.Projects, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withUsers; query != nil {
		if err := cq.loadUsers(ctx, query, nodes,
			func(n *Comment) { n.Edges.Users = []*User{} },
			func(n *Comment, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedEpics {
		if err := cq.loadEpics(ctx, query, nodes,
			func(n *Comment) { n.appendNamedEpics(name) },
			func(n *Comment, e *Epic) { n.appendNamedEpics(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedProjects {
		if err := cq.loadProjects(ctx, query, nodes,
			func(n *Comment) { n.appendNamedProjects(name) },
			func(n *Comment, e *Project) { n.appendNamedProjects(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedUsers {
		if err := cq.loadUsers(ctx, query, nodes,
			func(n *Comment) { n.appendNamedUsers(name) },
			func(n *Comment, e *User) { n.appendNamedUsers(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CommentQuery) loadEpics(ctx context.Context, query *EpicQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *Epic)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Comment)
	nids := make(map[int]map[*Comment]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(comment.EpicsTable)
		s.Join(joinT).On(s.C(epic.FieldID), joinT.C(comment.EpicsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(comment.EpicsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(comment.EpicsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Comment]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "epics" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CommentQuery) loadProjects(ctx context.Context, query *ProjectQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *Project)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Comment)
	nids := make(map[int]map[*Comment]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(comment.ProjectsTable)
		s.Join(joinT).On(s.C(project.FieldID), joinT.C(comment.ProjectsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(comment.ProjectsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(comment.ProjectsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Comment]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "projects" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CommentQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Comment)
	nids := make(map[int]map[*Comment]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(comment.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(comment.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(comment.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(comment.UsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Comment]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *CommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for i := range fields {
			if fields[i] != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(comment.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = comment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedEpics tells the query-builder to eager-load the nodes that are connected to the "epics"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CommentQuery) WithNamedEpics(name string, opts ...func(*EpicQuery)) *CommentQuery {
	query := (&EpicClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedEpics == nil {
		cq.withNamedEpics = make(map[string]*EpicQuery)
	}
	cq.withNamedEpics[name] = query
	return cq
}

// WithNamedProjects tells the query-builder to eager-load the nodes that are connected to the "projects"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CommentQuery) WithNamedProjects(name string, opts ...func(*ProjectQuery)) *CommentQuery {
	query := (&ProjectClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedProjects == nil {
		cq.withNamedProjects = make(map[string]*ProjectQuery)
	}
	cq.withNamedProjects[name] = query
	return cq
}

// WithNamedUsers tells the query-builder to eager-load the nodes that are connected to the "users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CommentQuery) WithNamedUsers(name string, opts ...func(*UserQuery)) *CommentQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedUsers == nil {
		cq.withNamedUsers = make(map[string]*UserQuery)
	}
	cq.withNamedUsers[name] = query
	return cq
}

// CommentGroupBy is the group-by builder for Comment entities.
type CommentGroupBy struct {
	selector
	build *CommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommentGroupBy) Aggregate(fns ...AggregateFunc) *CommentGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeComment, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CommentGroupBy) sqlScan(ctx context.Context, root *CommentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommentSelect is the builder for selecting fields of Comment entities.
type CommentSelect struct {
	*CommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CommentSelect) Aggregate(fns ...AggregateFunc) *CommentSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommentSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeComment, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentSelect](ctx, cs.CommentQuery, cs, cs.inters, v)
}

func (cs *CommentSelect) sqlScan(ctx context.Context, root *CommentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
