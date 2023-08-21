package migrations

type Migration interface {
    Up() error
    Down() error
}


func Run(face string) {
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
