package services

import (
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/repositories"
)

type AgencyService struct {
	Repo *repositories.AgencyRepository
}

func NewAgencyService(repo *repositories.AgencyRepository) *AgencyService {
	return &AgencyService{Repo: repo}
}

func (s *AgencyService) CreateAgency(a *models.Agency) error {
	return s.Repo.CreateAgency(a)
}

func (s *AgencyService) GetAgencyByClientID(id uint, search string) (*[]models.AgencyClientJoin, error) {
	return s.Repo.GetAgencyByClientID(id, search)
}
