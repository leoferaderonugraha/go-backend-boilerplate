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

func (db *DbConnection) Connect() (*gorm.DB, error) {
    url := config.Get[string]("database_url", "")

    if url == "" {
        panic("Database URL is not set")
    }

    dbConn, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        return nil, err
    }

    db.Db = dbConn

    return db.Db, nil
}
