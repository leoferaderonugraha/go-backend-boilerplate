package repositories

import (
	"leoferaderonugraha/go-backend-boilerplate/src/app/models"
    e "leoferaderonugraha/go-backend-boilerplate/pkg/errors"

	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(user *models.User) error {
    result := r.db.Create(user)

	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return e.USER_NOT_FOUND
        }

        return result.Error
	}

	return nil
}

func (r *UserRepository) GetUserById(id string) (*models.User, error) {
    user := &models.User{}
    result := r.db.Where("id = ?", id).First(user)

    if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, result.Error
    }

    return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := r.db.Where("email = ?", email).First(user)

	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, e.USER_NOT_FOUND
        }

        return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) Update(user *models.User) error {
    result := r.db.Save(user)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return e.USER_NOT_FOUND
        }

        return result.Error
    }

    return nil
}
