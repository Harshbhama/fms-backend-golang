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

func (s *FreelancerService) CreateFreelancerTimesheet(ft *models.FreelancerTimesheet) error {
	return s.Repo.CreateFreelancerTimesheet(ft)
}

func (s *FreelancerService) CreateFreelancerTimesheetMetadata(ftm *models.FreelancerTimesheetMetadata) error {
	return s.Repo.CreateFreelancerTimesheetMetadata(ftm)
}
func (s *FreelancerService) GetFreelancerByClientID(id uint, search string) (*[]models.FreelancerClientJoin, error) {
	return s.Repo.GetFreelancerByClientID(id, search)
}

func (s *FreelancerService) GetProjectsByFreelancerID(id uint, search string) (*[]models.FreelancerProjectJoin, error) {
	return s.Repo.GetProjectsByFreelancerID(id, search)
}

func (s *FreelancerService) GetClientsByFreelancerID(id uint, search string) (*[]models.ClientFreelancerJoin, error) {
	return s.Repo.GetClientsByFreelancerID(id, search)
}
