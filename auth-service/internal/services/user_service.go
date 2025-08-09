package services

import (
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}
 
func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

func (s *UserService) GetUser(id int64) (*models.User, error) {
	return s.userRepository.GetUser(id)
} 

func (s *UserService) UpdateUser(id int64, user *models.User) error {
	return s.userRepository.UpdateUser(id, user)
}

func (s *UserService) LoginUser(user *models.UserLogin) (*models.User, error) {
	return s.userRepository.LoginUser(user)
}