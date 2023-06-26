package database

import (
    "gorm.io/gorm"
)

func WithTx(db *gorm.DB, fn func (db *gorm.DB) error) {
    tx := db.Begin()
    if err := fn(tx); err != nil {
        tx.Rollback()
        panic(err)
    }
    tx.Commit()
}

