package database

import (
    "leoferaderonugraha/go-backend-boilerplate/pkg/config"

    "gorm.io/gorm"
)

func WithTx(cfg *config.Config, fn func (db *gorm.DB) error) {
    conn := NewDb()
    conn.Connect(cfg)

    tx := conn.Db.Begin()
    if err := fn(tx); err != nil {
        tx.Rollback()
        panic(err)
    }
    tx.Commit()
}

