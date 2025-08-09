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

func (s *ClientService) CreateClientFreelancer(clientFreelancer *models.ClientFreelancer) error {
	// This function is not implemented yet
	// Here you would typically call a repository method to create the client-freelancer relationship
	// For now, we will just return nil
	return s.clientRepository.CreateClientFreelancer(clientFreelancer)
}
