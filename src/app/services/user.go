package services

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/models"
    "leoferaderonugraha/go-backend-boilerplate/src/app/repositories"

	"errors"
)

type UserRegistrationService struct {
	userRepository repositories.UserRepository
}

func NewUserRegistrationService(userRepository repositories.UserRepository) *UserRegistrationService {
	return &UserRegistrationService{
		userRepository: userRepository,
	}
}

func (s *UserRegistrationService) RegisterUser(name, email, password string) (*models.User, error) {
	existingUser, _ := s.userRepository.GetUserByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user with the given email already exists")
	}

	user := &models.User{
		Name:      name,
		Email:     email,
		Password:  password,
	}

	err := s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

