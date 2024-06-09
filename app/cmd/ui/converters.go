package ui

import "github.com/jenmud/consensus/business/service"

// CoreProjectToProject converts a core project to a project.
func CoreProjectToProject(project *service.Project) Project {
	return Project{
		ID:          project.GetId(),
		Title:       project.GetName(),
		Description: project.GetDescription(),
		Owner:       CoreUserToUser(project.GetOwner()),
	}
}

// CoreProjectsToProjects converts a one or more core projects to one or more projects.
func CoreProjectsToProjects(projects ...*service.Project) []Project {
	projectsCore := make([]Project, len(projects))
	for i, project := range projects {
		projectsCore[i] = CoreProjectToProject(project)
	}
	return projectsCore
}

// CoreUserToUser converts a core user to a user.
func CoreUserToUser(user *service.User) User {
	return User{
		ID:        user.GetId(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
		Role:      user.GetRole().String(),
	}
}

// CoreUsersToUsers converts a one or more core users to one or more users.
func CoreUsersToUsers(users ...*service.User) []User {
	usersCore := make([]User, len(users))
	for i, user := range users {
		usersCore[i] = CoreUserToUser(user)
	}
	return usersCore
}
