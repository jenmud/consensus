// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/jenmud/consensus/ent"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Comment struct {
		Epics    func(childComplexity int) int
		ID       func(childComplexity int) int
		Projects func(childComplexity int) int
		Text     func(childComplexity int) int
		Users    func(childComplexity int) int
	}

	CommentConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	CommentEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Epic struct {
		Assignee    func(childComplexity int) int
		Comments    func(childComplexity int) int
		Description func(childComplexity int) int
		ID          func(childComplexity int) int
		Name        func(childComplexity int) int
		Project     func(childComplexity int) int
		Reporter    func(childComplexity int) int
	}

	EpicConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	EpicEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Project struct {
		Comments    func(childComplexity int) int
		Description func(childComplexity int) int
		Epics       func(childComplexity int) int
		ID          func(childComplexity int) int
		Name        func(childComplexity int) int
		Owner       func(childComplexity int) int
	}

	ProjectConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	ProjectEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Query struct {
		Comments func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.CommentWhereInput) int
		Epics    func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.EpicWhereInput) int
		Node     func(childComplexity int, id int) int
		Nodes    func(childComplexity int, ids []int) int
		Projects func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ProjectWhereInput) int
		Users    func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) int
	}

	User struct {
		Assignee func(childComplexity int) int
		Comments func(childComplexity int) int
		Email    func(childComplexity int) int
		ID       func(childComplexity int) int
		Name     func(childComplexity int) int
		Owns     func(childComplexity int) int
		Reporter func(childComplexity int) int
		Surname  func(childComplexity int) int
		Username func(childComplexity int) int
	}

	UserConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	UserEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Comment.epics":
		if e.complexity.Comment.Epics == nil {
			break
		}

		return e.complexity.Comment.Epics(childComplexity), true

	case "Comment.id":
		if e.complexity.Comment.ID == nil {
			break
		}

		return e.complexity.Comment.ID(childComplexity), true

	case "Comment.projects":
		if e.complexity.Comment.Projects == nil {
			break
		}

		return e.complexity.Comment.Projects(childComplexity), true

	case "Comment.text":
		if e.complexity.Comment.Text == nil {
			break
		}

		return e.complexity.Comment.Text(childComplexity), true

	case "Comment.users":
		if e.complexity.Comment.Users == nil {
			break
		}

		return e.complexity.Comment.Users(childComplexity), true

	case "CommentConnection.edges":
		if e.complexity.CommentConnection.Edges == nil {
			break
		}

		return e.complexity.CommentConnection.Edges(childComplexity), true

	case "CommentConnection.pageInfo":
		if e.complexity.CommentConnection.PageInfo == nil {
			break
		}

		return e.complexity.CommentConnection.PageInfo(childComplexity), true

	case "CommentConnection.totalCount":
		if e.complexity.CommentConnection.TotalCount == nil {
			break
		}

		return e.complexity.CommentConnection.TotalCount(childComplexity), true

	case "CommentEdge.cursor":
		if e.complexity.CommentEdge.Cursor == nil {
			break
		}

		return e.complexity.CommentEdge.Cursor(childComplexity), true

	case "CommentEdge.node":
		if e.complexity.CommentEdge.Node == nil {
			break
		}

		return e.complexity.CommentEdge.Node(childComplexity), true

	case "Epic.assignee":
		if e.complexity.Epic.Assignee == nil {
			break
		}

		return e.complexity.Epic.Assignee(childComplexity), true

	case "Epic.comments":
		if e.complexity.Epic.Comments == nil {
			break
		}

		return e.complexity.Epic.Comments(childComplexity), true

	case "Epic.description":
		if e.complexity.Epic.Description == nil {
			break
		}

		return e.complexity.Epic.Description(childComplexity), true

	case "Epic.id":
		if e.complexity.Epic.ID == nil {
			break
		}

		return e.complexity.Epic.ID(childComplexity), true

	case "Epic.name":
		if e.complexity.Epic.Name == nil {
			break
		}

		return e.complexity.Epic.Name(childComplexity), true

	case "Epic.project":
		if e.complexity.Epic.Project == nil {
			break
		}

		return e.complexity.Epic.Project(childComplexity), true

	case "Epic.reporter":
		if e.complexity.Epic.Reporter == nil {
			break
		}

		return e.complexity.Epic.Reporter(childComplexity), true

	case "EpicConnection.edges":
		if e.complexity.EpicConnection.Edges == nil {
			break
		}

		return e.complexity.EpicConnection.Edges(childComplexity), true

	case "EpicConnection.pageInfo":
		if e.complexity.EpicConnection.PageInfo == nil {
			break
		}

		return e.complexity.EpicConnection.PageInfo(childComplexity), true

	case "EpicConnection.totalCount":
		if e.complexity.EpicConnection.TotalCount == nil {
			break
		}

		return e.complexity.EpicConnection.TotalCount(childComplexity), true

	case "EpicEdge.cursor":
		if e.complexity.EpicEdge.Cursor == nil {
			break
		}

		return e.complexity.EpicEdge.Cursor(childComplexity), true

	case "EpicEdge.node":
		if e.complexity.EpicEdge.Node == nil {
			break
		}

		return e.complexity.EpicEdge.Node(childComplexity), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Project.comments":
		if e.complexity.Project.Comments == nil {
			break
		}

		return e.complexity.Project.Comments(childComplexity), true

	case "Project.description":
		if e.complexity.Project.Description == nil {
			break
		}

		return e.complexity.Project.Description(childComplexity), true

	case "Project.epics":
		if e.complexity.Project.Epics == nil {
			break
		}

		return e.complexity.Project.Epics(childComplexity), true

	case "Project.id":
		if e.complexity.Project.ID == nil {
			break
		}

		return e.complexity.Project.ID(childComplexity), true

	case "Project.name":
		if e.complexity.Project.Name == nil {
			break
		}

		return e.complexity.Project.Name(childComplexity), true

	case "Project.owner":
		if e.complexity.Project.Owner == nil {
			break
		}

		return e.complexity.Project.Owner(childComplexity), true

	case "ProjectConnection.edges":
		if e.complexity.ProjectConnection.Edges == nil {
			break
		}

		return e.complexity.ProjectConnection.Edges(childComplexity), true

	case "ProjectConnection.pageInfo":
		if e.complexity.ProjectConnection.PageInfo == nil {
			break
		}

		return e.complexity.ProjectConnection.PageInfo(childComplexity), true

	case "ProjectConnection.totalCount":
		if e.complexity.ProjectConnection.TotalCount == nil {
			break
		}

		return e.complexity.ProjectConnection.TotalCount(childComplexity), true

	case "ProjectEdge.cursor":
		if e.complexity.ProjectEdge.Cursor == nil {
			break
		}

		return e.complexity.ProjectEdge.Cursor(childComplexity), true

	case "ProjectEdge.node":
		if e.complexity.ProjectEdge.Node == nil {
			break
		}

		return e.complexity.ProjectEdge.Node(childComplexity), true

	case "Query.comments":
		if e.complexity.Query.Comments == nil {
			break
		}

		args, err := ec.field_Query_comments_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Comments(childComplexity, args["after"].(*ent.Cursor), args["first"].(*int), args["before"].(*ent.Cursor), args["last"].(*int), args["where"].(*ent.CommentWhereInput)), true

	case "Query.epics":
		if e.complexity.Query.Epics == nil {
			break
		}

		args, err := ec.field_Query_epics_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Epics(childComplexity, args["after"].(*ent.Cursor), args["first"].(*int), args["before"].(*ent.Cursor), args["last"].(*int), args["where"].(*ent.EpicWhereInput)), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(int)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]int)), true

	case "Query.projects":
		if e.complexity.Query.Projects == nil {
			break
		}

		args, err := ec.field_Query_projects_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Projects(childComplexity, args["after"].(*ent.Cursor), args["first"].(*int), args["before"].(*ent.Cursor), args["last"].(*int), args["where"].(*ent.ProjectWhereInput)), true

	case "Query.users":
		if e.complexity.Query.Users == nil {
			break
		}

		args, err := ec.field_Query_users_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Users(childComplexity, args["after"].(*ent.Cursor), args["first"].(*int), args["before"].(*ent.Cursor), args["last"].(*int), args["where"].(*ent.UserWhereInput)), true

	case "User.assignee":
		if e.complexity.User.Assignee == nil {
			break
		}

		return e.complexity.User.Assignee(childComplexity), true

	case "User.comments":
		if e.complexity.User.Comments == nil {
			break
		}

		return e.complexity.User.Comments(childComplexity), true

	case "User.email":
		if e.complexity.User.Email == nil {
			break
		}

		return e.complexity.User.Email(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "User.owns":
		if e.complexity.User.Owns == nil {
			break
		}

		return e.complexity.User.Owns(childComplexity), true

	case "User.reporter":
		if e.complexity.User.Reporter == nil {
			break
		}

		return e.complexity.User.Reporter(childComplexity), true

	case "User.surname":
		if e.complexity.User.Surname == nil {
			break
		}

		return e.complexity.User.Surname(childComplexity), true

	case "User.username":
		if e.complexity.User.Username == nil {
			break
		}

		return e.complexity.User.Username(childComplexity), true

	case "UserConnection.edges":
		if e.complexity.UserConnection.Edges == nil {
			break
		}

		return e.complexity.UserConnection.Edges(childComplexity), true

	case "UserConnection.pageInfo":
		if e.complexity.UserConnection.PageInfo == nil {
			break
		}

		return e.complexity.UserConnection.PageInfo(childComplexity), true

	case "UserConnection.totalCount":
		if e.complexity.UserConnection.TotalCount == nil {
			break
		}

		return e.complexity.UserConnection.TotalCount(childComplexity), true

	case "UserEdge.cursor":
		if e.complexity.UserEdge.Cursor == nil {
			break
		}

		return e.complexity.UserEdge.Cursor(childComplexity), true

	case "UserEdge.node":
		if e.complexity.UserEdge.Node == nil {
			break
		}

		return e.complexity.UserEdge.Node(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCommentWhereInput,
		ec.unmarshalInputCreateCommentInput,
		ec.unmarshalInputCreateEpicInput,
		ec.unmarshalInputCreateProjectInput,
		ec.unmarshalInputCreateUserInput,
		ec.unmarshalInputEpicWhereInput,
		ec.unmarshalInputProjectWhereInput,
		ec.unmarshalInputUpdateCommentInput,
		ec.unmarshalInputUpdateEpicInput,
		ec.unmarshalInputUpdateProjectInput,
		ec.unmarshalInputUpdateUserInput,
		ec.unmarshalInputUserWhereInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../ent.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
type Comment implements Node {
  id: ID!
  text: String!
  epics: [Epic!]
  projects: [Project!]
  users: [User!]
}
"""A connection to a list of items."""
type CommentConnection {
  """A list of edges."""
  edges: [CommentEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type CommentEdge {
  """The item at the end of the edge."""
  node: Comment
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
CommentWhereInput is used for filtering Comment objects.
Input was generated by ent.
"""
input CommentWhereInput {
  not: CommentWhereInput
  and: [CommentWhereInput!]
  or: [CommentWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """text field predicates"""
  text: String
  textNEQ: String
  textIn: [String!]
  textNotIn: [String!]
  textGT: String
  textGTE: String
  textLT: String
  textLTE: String
  textContains: String
  textHasPrefix: String
  textHasSuffix: String
  textEqualFold: String
  textContainsFold: String
  """epics edge predicates"""
  hasEpics: Boolean
  hasEpicsWith: [EpicWhereInput!]
  """projects edge predicates"""
  hasProjects: Boolean
  hasProjectsWith: [ProjectWhereInput!]
  """users edge predicates"""
  hasUsers: Boolean
  hasUsersWith: [UserWhereInput!]
}
"""
CreateCommentInput is used for create Comment object.
Input was generated by ent.
"""
input CreateCommentInput {
  text: String!
  epicIDs: [ID!]
  projectIDs: [ID!]
  userIDs: [ID!]
}
"""
CreateEpicInput is used for create Epic object.
Input was generated by ent.
"""
input CreateEpicInput {
  name: String!
  description: String!
  projectID: ID
  reporterID: ID
  assigneeID: ID
  commentIDs: [ID!]
}
"""
CreateProjectInput is used for create Project object.
Input was generated by ent.
"""
input CreateProjectInput {
  name: String!
  description: String
  epicIDs: [ID!]
  ownerID: ID
  commentIDs: [ID!]
}
"""
CreateUserInput is used for create User object.
Input was generated by ent.
"""
input CreateUserInput {
  name: String!
  surname: String!
  username: String!
  email: String!
  ownIDs: [ID!]
  reporterIDs: [ID!]
  assigneeIDs: [ID!]
  commentIDs: [ID!]
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
type Epic implements Node {
  id: ID!
  name: String!
  description: String!
  project: Project
  reporter: User
  assignee: User
  comments: [Comment!]
}
"""A connection to a list of items."""
type EpicConnection {
  """A list of edges."""
  edges: [EpicEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type EpicEdge {
  """The item at the end of the edge."""
  node: Epic
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
EpicWhereInput is used for filtering Epic objects.
Input was generated by ent.
"""
input EpicWhereInput {
  not: EpicWhereInput
  and: [EpicWhereInput!]
  or: [EpicWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """description field predicates"""
  description: String
  descriptionNEQ: String
  descriptionIn: [String!]
  descriptionNotIn: [String!]
  descriptionGT: String
  descriptionGTE: String
  descriptionLT: String
  descriptionLTE: String
  descriptionContains: String
  descriptionHasPrefix: String
  descriptionHasSuffix: String
  descriptionEqualFold: String
  descriptionContainsFold: String
  """project edge predicates"""
  hasProject: Boolean
  hasProjectWith: [ProjectWhereInput!]
  """reporter edge predicates"""
  hasReporter: Boolean
  hasReporterWith: [UserWhereInput!]
  """assignee edge predicates"""
  hasAssignee: Boolean
  hasAssigneeWith: [UserWhereInput!]
  """comments edge predicates"""
  hasComments: Boolean
  hasCommentsWith: [CommentWhereInput!]
}
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "github.com/jenmud/consensus/ent.Noder") {
  """The id of the object."""
  id: ID!
}
"""Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument."""
enum OrderDirection {
  """Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  ASC
  """Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
type Project implements Node {
  id: ID!
  name: String!
  description: String
  epics: [Epic!]
  owner: User
  comments: [Comment!]
}
"""A connection to a list of items."""
type ProjectConnection {
  """A list of edges."""
  edges: [ProjectEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type ProjectEdge {
  """The item at the end of the edge."""
  node: Project
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
ProjectWhereInput is used for filtering Project objects.
Input was generated by ent.
"""
input ProjectWhereInput {
  not: ProjectWhereInput
  and: [ProjectWhereInput!]
  or: [ProjectWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """description field predicates"""
  description: String
  descriptionNEQ: String
  descriptionIn: [String!]
  descriptionNotIn: [String!]
  descriptionGT: String
  descriptionGTE: String
  descriptionLT: String
  descriptionLTE: String
  descriptionContains: String
  descriptionHasPrefix: String
  descriptionHasSuffix: String
  descriptionIsNil: Boolean
  descriptionNotNil: Boolean
  descriptionEqualFold: String
  descriptionContainsFold: String
  """epics edge predicates"""
  hasEpics: Boolean
  hasEpicsWith: [EpicWhereInput!]
  """owner edge predicates"""
  hasOwner: Boolean
  hasOwnerWith: [UserWhereInput!]
  """comments edge predicates"""
  hasComments: Boolean
  hasCommentsWith: [CommentWhereInput!]
}
type Query {
  """Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node
  """Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!
  comments(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for Comments returned from the connection."""
    where: CommentWhereInput
  ): CommentConnection!
  epics(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for Epics returned from the connection."""
    where: EpicWhereInput
  ): EpicConnection!
  projects(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for Projects returned from the connection."""
    where: ProjectWhereInput
  ): ProjectConnection!
  users(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for Users returned from the connection."""
    where: UserWhereInput
  ): UserConnection!
}
"""
UpdateCommentInput is used for update Comment object.
Input was generated by ent.
"""
input UpdateCommentInput {
  text: String
  addEpicIDs: [ID!]
  removeEpicIDs: [ID!]
  addProjectIDs: [ID!]
  removeProjectIDs: [ID!]
  addUserIDs: [ID!]
  removeUserIDs: [ID!]
}
"""
UpdateEpicInput is used for update Epic object.
Input was generated by ent.
"""
input UpdateEpicInput {
  name: String
  description: String
  clearProject: Boolean
  projectID: ID
  clearReporter: Boolean
  reporterID: ID
  clearAssignee: Boolean
  assigneeID: ID
  addCommentIDs: [ID!]
  removeCommentIDs: [ID!]
}
"""
UpdateProjectInput is used for update Project object.
Input was generated by ent.
"""
input UpdateProjectInput {
  name: String
  clearDescription: Boolean
  description: String
  addEpicIDs: [ID!]
  removeEpicIDs: [ID!]
  clearOwner: Boolean
  ownerID: ID
  addCommentIDs: [ID!]
  removeCommentIDs: [ID!]
}
"""
UpdateUserInput is used for update User object.
Input was generated by ent.
"""
input UpdateUserInput {
  name: String
  surname: String
  username: String
  email: String
  addOwnIDs: [ID!]
  removeOwnIDs: [ID!]
  addReporterIDs: [ID!]
  removeReporterIDs: [ID!]
  addAssigneeIDs: [ID!]
  removeAssigneeIDs: [ID!]
  addCommentIDs: [ID!]
  removeCommentIDs: [ID!]
}
type User implements Node {
  id: ID!
  name: String!
  surname: String!
  username: String!
  email: String!
  owns: [Project!]
  reporter: [Epic!]
  assignee: [Epic!]
  comments: [Comment!]
}
"""A connection to a list of items."""
type UserConnection {
  """A list of edges."""
  edges: [UserEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type UserEdge {
  """The item at the end of the edge."""
  node: User
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
UserWhereInput is used for filtering User objects.
Input was generated by ent.
"""
input UserWhereInput {
  not: UserWhereInput
  and: [UserWhereInput!]
  or: [UserWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """surname field predicates"""
  surname: String
  surnameNEQ: String
  surnameIn: [String!]
  surnameNotIn: [String!]
  surnameGT: String
  surnameGTE: String
  surnameLT: String
  surnameLTE: String
  surnameContains: String
  surnameHasPrefix: String
  surnameHasSuffix: String
  surnameEqualFold: String
  surnameContainsFold: String
  """username field predicates"""
  username: String
  usernameNEQ: String
  usernameIn: [String!]
  usernameNotIn: [String!]
  usernameGT: String
  usernameGTE: String
  usernameLT: String
  usernameLTE: String
  usernameContains: String
  usernameHasPrefix: String
  usernameHasSuffix: String
  usernameEqualFold: String
  usernameContainsFold: String
  """email field predicates"""
  email: String
  emailNEQ: String
  emailIn: [String!]
  emailNotIn: [String!]
  emailGT: String
  emailGTE: String
  emailLT: String
  emailLTE: String
  emailContains: String
  emailHasPrefix: String
  emailHasSuffix: String
  emailEqualFold: String
  emailContainsFold: String
  """owns edge predicates"""
  hasOwns: Boolean
  hasOwnsWith: [ProjectWhereInput!]
  """reporter edge predicates"""
  hasReporter: Boolean
  hasReporterWith: [EpicWhereInput!]
  """assignee edge predicates"""
  hasAssignee: Boolean
  hasAssigneeWith: [EpicWhereInput!]
  """comments edge predicates"""
  hasComments: Boolean
  hasCommentsWith: [CommentWhereInput!]
}
`, BuiltIn: false},
	{Name: "../scalars.graphql", Input: `scalar Time
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
