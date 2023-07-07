package services

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/models"
    "leoferaderonugraha/go-backend-boilerplate/src/app/repositories"
    e "leoferaderonugraha/go-backend-boilerplate/pkg/errors"

)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserRegistrationService(repository repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

type UserRegistrationRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserDetailRequest struct {
    Id string `json:"id"`
}

func (s *UserService) Register(name, email, password string) (*models.User, error) {
	existingUser, _ := s.repository.GetUserByEmail(email)

	if existingUser != nil {
		return nil, e.USER_ALREADY_EXISTS
	}

	user := &models.User{
		Name:      name,
		Email:     email,
		Password:  password,
	}

	err := s.repository.Save(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Details(id string) (*models.User, error) {
    user, err := s.repository.GetUserById(id)

    if err != nil {
        return nil, err
    }

    return user, nil
}

func (s *UserService) Update(id string, data map[string]any) (*models.User, error) {
    user, err := s.repository.GetUserById(id)

    if err != nil {
        return nil, err
    }

    if data["name"] != nil {
        user.Name = data["name"].(string)
    }
    
    if data["email"] != nil {
        user.Email = data["email"].(string)
    }

    if data["password"] != nil {
        user.Password = data["password"].(string)
    }

    err = s.repository.Update(user)

    if err != nil {
        return nil, err
    }

    return user, nil
}
