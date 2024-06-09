package ui

import "github.com/jenmud/consensus/business/service"

// CoreProjectToProject converts a core project to a project.
func CoreProjectToProject(project *service.Project) Project {
	return Project{
		ID:          project.GetId(),
		Title:       project.GetName(),
		Description: project.GetDescription(),
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
