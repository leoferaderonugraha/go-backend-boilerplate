package migrations

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/models"

    "gorm.io/gorm"
)

type DropUserDeletedAt struct {
    Db *gorm.DB
}

func NewDropUserDeletedAt(db *gorm.DB) *DropUserDeletedAt {
    return &DropUserDeletedAt{
        Db: db,
    }
}

func (m *DropUserDeletedAt) Up() error {
    m.Db.Migrator().DropColumn(&models.User{}, "deleted_at")

    return m.Db.Error
}

func (m *DropUserDeletedAt) Down() error {
    return nil
}
