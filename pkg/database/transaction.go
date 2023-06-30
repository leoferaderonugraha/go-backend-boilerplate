package database

import (
    "gorm.io/gorm"
)

func WithTx(fn func (db *gorm.DB) error) {
    conn := New()

    tx := conn.Db.Begin()
    if err := fn(tx); err != nil {
        tx.Rollback()
        panic(err)
    }
    tx.Commit()
}

