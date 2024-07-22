package service

import (
	context "context"
	"database/sql"
	"log/slog"

	"github.com/jenmud/consensus/foundation/crypto"
	client "github.com/jenmud/consensus/foundation/data/sqlite"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Consensus is the main consensus service.
type Service struct {
	UnimplementedConsensusServer
	db     *sql.DB
	client *client.Queries
}

// New creates a new service.
func New(db *sql.DB) *Service {
	return &Service{
		db:     db, // we need the actual db because we need to create transactions.
		client: client.New(db),
	}
}

// AuthenticateUser authenticates a user and returns the user if valid.
func (s *Service) AuthenticateUser(ctx context.Context, req *AuthReq) (*User, error) {
	user, err := s.client.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		slog.Error("failed to get user", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.NotFound, "failed to get user: %v", err)
	}

	if !crypto.CheckPasswordHash(user.Password, req.GetPassword()) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	return DBUserToCoreUser(user), status.Errorf(codes.OK, "")
}

// CreateUser creates a new user.
func (s *Service) CreateUser(ctx context.Context, user *User) (*User, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("failed to begin transaction", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}

	defer tx.Rollback()

	// we do not want to store the password in plain text, so we hash it using the bcrypt algorithm.
	hashedPassword, err := crypto.HashPassword(user.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err.Error())
	}

	c := s.client.WithTx(tx)
	u, err := c.CreateUser(
		ctx,
		client.CreateUserParams{
			Email:     user.GetEmail(),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
			Password:  hashedPassword,
			Role:      user.GetRole().String(),
		},
	)

	if err != nil {
		slog.Error("failed to create user", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	if err := tx.Commit(); err != nil {
		slog.Error("failed to commit transaction", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

	return DBUserToCoreUser(u), status.Error(codes.OK, "")
}

// GetUsers returns all the existing users.
func (s *Service) GetUsers(ctx context.Context, req *GetUsersReq) (*Users, error) {
	users, err := s.client.GetUsers(ctx)
	if err != nil {
		slog.Error("failed to get users", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to get users: %v", err)
	}

	return &Users{Users: DBUsersToCoreUsers(users...)}, status.Error(codes.OK, "")
}

// CreateProject creates a new project.
func (s *Service) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("failed to begin transaction", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}

	defer tx.Rollback()

	c := s.client.WithTx(tx)
	p, err := c.CreateProject(
		ctx,
		client.CreateProjectParams{
			Name: project.GetName(),
			Description: sql.NullString{
				String: project.GetDescription(),
				Valid:  project.GetDescription() != "",
			},
			UserID: project.GetOwner().GetId(),
		},
	)

	if err != nil {
		slog.Error("failed to create project", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to create project: %v", err)
	}

	if err := tx.Commit(); err != nil {
		slog.Error("failed to commit transaction", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

	return DBProjectToCoreProject(p), status.Error(codes.OK, "")
}

// GetProjects returns all the existing projects.
func (s *Service) GetProjects(ctx context.Context, req *ProjectsReq) (*Projects, error) {
	rows, err := s.client.GetProjects(ctx)
	if err != nil {
		slog.Error("failed to get projects", slog.String("reason", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to get projects: %v", err)
	}

	projects := &Projects{Projects: make([]*Project, len(rows))}
	for i, row := range rows {
		p := DBProjectToCoreProject(row.Project)
		p.Owner = DBUserToCoreUser(row.User)
		projects.Projects[i] = p
	}

	return projects, status.Error(codes.OK, "")
}
