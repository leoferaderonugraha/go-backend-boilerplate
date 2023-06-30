package database

import (
    "leoferaderonugraha/go-backend-boilerplate/pkg/config"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

type DbConnection struct {
    Db *gorm.DB
}

func NewDb() *DbConnection {
    return &DbConnection{}
}

func (db *DbConnection) Connect(config *config.Config) (*gorm.DB, error) {
    dbConn, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})

    if err != nil {
        return nil, err
    }

    db.Db = dbConn

    return db.Db, nil
}
