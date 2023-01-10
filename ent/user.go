// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/jenmud/consensus/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Surname holds the value of the "surname" field.
	Surname string `json:"surname,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Owns holds the value of the owns edge.
	Owns []*Project `json:"owns,omitempty"`
	// Reporter holds the value of the reporter edge.
	Reporter []*Epic `json:"reporter,omitempty"`
	// Assignee holds the value of the assignee edge.
	Assignee []*Epic `json:"assignee,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedOwns     map[string][]*Project
	namedReporter map[string][]*Epic
	namedAssignee map[string][]*Epic
	namedComments map[string][]*Comment
}

// OwnsOrErr returns the Owns value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OwnsOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Owns, nil
	}
	return nil, &NotLoadedError{edge: "owns"}
}

// ReporterOrErr returns the Reporter value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReporterOrErr() ([]*Epic, error) {
	if e.loadedTypes[1] {
		return e.Reporter, nil
	}
	return nil, &NotLoadedError{edge: "reporter"}
}

// AssigneeOrErr returns the Assignee value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AssigneeOrErr() ([]*Epic, error) {
	if e.loadedTypes[2] {
		return e.Assignee, nil
	}
	return nil, &NotLoadedError{edge: "assignee"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[3] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldSurname, user.FieldUsername, user.FieldEmail:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldSurname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field surname", values[i])
			} else if value.Valid {
				u.Surname = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		}
	}
	return nil
}

// QueryOwns queries the "owns" edge of the User entity.
func (u *User) QueryOwns() *ProjectQuery {
	return (&UserClient{config: u.config}).QueryOwns(u)
}

// QueryReporter queries the "reporter" edge of the User entity.
func (u *User) QueryReporter() *EpicQuery {
	return (&UserClient{config: u.config}).QueryReporter(u)
}

// QueryAssignee queries the "assignee" edge of the User entity.
func (u *User) QueryAssignee() *EpicQuery {
	return (&UserClient{config: u.config}).QueryAssignee(u)
}

// QueryComments queries the "comments" edge of the User entity.
func (u *User) QueryComments() *CommentQuery {
	return (&UserClient{config: u.config}).QueryComments(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("surname=")
	builder.WriteString(u.Surname)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteByte(')')
	return builder.String()
}

// NamedOwns returns the Owns named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOwns(name string) ([]*Project, error) {
	if u.Edges.namedOwns == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOwns[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOwns(name string, edges ...*Project) {
	if u.Edges.namedOwns == nil {
		u.Edges.namedOwns = make(map[string][]*Project)
	}
	if len(edges) == 0 {
		u.Edges.namedOwns[name] = []*Project{}
	} else {
		u.Edges.namedOwns[name] = append(u.Edges.namedOwns[name], edges...)
	}
}

// NamedReporter returns the Reporter named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedReporter(name string) ([]*Epic, error) {
	if u.Edges.namedReporter == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedReporter[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedReporter(name string, edges ...*Epic) {
	if u.Edges.namedReporter == nil {
		u.Edges.namedReporter = make(map[string][]*Epic)
	}
	if len(edges) == 0 {
		u.Edges.namedReporter[name] = []*Epic{}
	} else {
		u.Edges.namedReporter[name] = append(u.Edges.namedReporter[name], edges...)
	}
}

// NamedAssignee returns the Assignee named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAssignee(name string) ([]*Epic, error) {
	if u.Edges.namedAssignee == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAssignee[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAssignee(name string, edges ...*Epic) {
	if u.Edges.namedAssignee == nil {
		u.Edges.namedAssignee = make(map[string][]*Epic)
	}
	if len(edges) == 0 {
		u.Edges.namedAssignee[name] = []*Epic{}
	} else {
		u.Edges.namedAssignee[name] = append(u.Edges.namedAssignee[name], edges...)
	}
}

// NamedComments returns the Comments named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedComments(name string) ([]*Comment, error) {
	if u.Edges.namedComments == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedComments[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedComments(name string, edges ...*Comment) {
	if u.Edges.namedComments == nil {
		u.Edges.namedComments = make(map[string][]*Comment)
	}
	if len(edges) == 0 {
		u.Edges.namedComments[name] = []*Comment{}
	} else {
		u.Edges.namedComments[name] = append(u.Edges.namedComments[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
