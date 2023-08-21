package migrations

import (
    "leoferaderonugraha/go-backend-boilerplate/pkg/config"
    "leoferaderonugraha/go-backend-boilerplate/pkg/database"
)

type Migration interface {
    Up() error
    Down() error
}


func Run(face string) {
    cfg, err := config.GetConfig()

    if err != nil {
        panic(err)
    }

    conn := database.NewDb()
    conn.Connect(cfg)

    if err != nil {
        panic(err)
    }

    listMigrations := make([]Migration, 0)
    // append list of migration here

    for _, migration := range listMigrations {
        var err error

        if face == "up" {
            err = migration.Up()
        } else {
            err = migration.Down()
        }

        if err != nil {
            panic(err)
        }
    }
}
