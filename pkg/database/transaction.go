package database

import (
    "gorm.io/gorm"
)

func WithTx(fn func (db *gorm.DB) error) {
    conn := NewDb()
    conn.Connect()

    tx := conn.Db.Begin()
    if err := fn(tx); err != nil {
        tx.Rollback()
    }
    tx.Commit()
}

