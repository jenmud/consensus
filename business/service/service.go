package service

import (
	context "context"
	"database/sql"

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

// CreateUser creates a new user.
func (s *Service) CreateUser(ctx context.Context, user *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// GetUsers returns all the existing users.
func (s *Service) GetUsers(ctx context.Context, req *GetUsersReq) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}

// CreateProject creates a new project.
func (s *Service) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}

// GetProjects returns all the existing projects.
func (s *Service) GetProjects(ctx context.Context, req *ProjectsReq) (*Projects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
