package service

import (
	client "github.com/jenmud/consensus/foundation/data/sqlite"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DBUserToCoreUser converts a database user to a core user.
func DBUserToCoreUser(user client.User) *User {
	var role Role

	switch user.Role {
	case "admin":
		role = Role_ADMIN
	case "user":
		role = Role_USER
	default:
		role = Role_USER
	}

	return &User{
		Id:        user.ID,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      role,
	}
}

// DBUsersToCoreUsers converts a one or more database users to one or more core users.
func DBUsersToCoreUsers(users ...client.User) []*User {
	usersCore := make([]*User, len(users))

	for i, user := range users {
		usersCore[i] = DBUserToCoreUser(user)
	}

	return usersCore
}

// DBProjectToCoreProject converts a database project to a core project.
func DBProjectToCoreProject(project client.Project) *Project {
	return &Project{
		Id:          project.ID,
		CreatedAt:   timestamppb.New(project.CreatedAt),
		UpdatedAt:   timestamppb.New(project.UpdatedAt),
		Name:        project.Name,
		Description: project.Description.String,
	}
}

// DBProjectsToCoreProjects converts a one or more database projects to one or more core projects.
func DBProjectsToCoreProjects(projects ...client.Project) []*Project {
	projectsCore := make([]*Project, len(projects))

	for i, project := range projects {
		projectsCore[i] = DBProjectToCoreProject(project)
	}

	return projectsCore
}
