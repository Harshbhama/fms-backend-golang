package services

import (
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/repositories"
)

type FreelancerService struct {
	Repo *repositories.FreelancerRepository
}

func NewFreelancerService(repo *repositories.FreelancerRepository) *FreelancerService {
	return &FreelancerService{Repo: repo}
}

func (s *FreelancerService) CreateFreelancer(f *models.Freelancer) error {
	return s.Repo.CreateFreelancer(f)
}

func (s *FreelancerService) CreateFreelancerRates(fr *models.FreelancerRates) error {
	return s.Repo.CreateFreelancerRates(fr)
}

func (s *FreelancerService) CreateFreelancerProject(fp *models.FreelancerProject) error {

	return s.Repo.CreateFreelancerProject(fp)
}