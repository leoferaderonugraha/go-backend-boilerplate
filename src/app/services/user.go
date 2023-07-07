package services

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/models"
    "leoferaderonugraha/go-backend-boilerplate/src/app/repositories"
    e "leoferaderonugraha/go-backend-boilerplate/pkg/errors"

)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserRegistrationService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

type UserRegistrationRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserRegistrationResponse struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func (s *UserService) Register(name, email, password string) (*models.User, error) {
	existingUser, _ := s.userRepository.GetUserByEmail(email)

	if existingUser != nil {
		return nil, e.USER_ALREADY_EXISTS
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

