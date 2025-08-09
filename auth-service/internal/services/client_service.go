package services

import (
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/repositories"
)

type ClientService struct {
	clientRepository *repositories.ClientRepository
}
 
func NewClientService(clientRepository *repositories.ClientRepository) *ClientService {
	return &ClientService{clientRepository: clientRepository}
}

func (s *ClientService) CreateClient(client *models.Client) error {
	return s.clientRepository.CreateClient(client)
}
