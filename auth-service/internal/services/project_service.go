package services

import (
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/repositories"
)

type ProjectService struct {
	Repo *repositories.ProjectRepository
}

func NewProjectService(repo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) CreateProject(p *models.Projects) error {
	return s.Repo.CreateProject(p)
}

func (s *ProjectService) SearchProjects(keyword string) ([]models.Projects, error) {
	return s.Repo.SearchProjects(keyword)
}
