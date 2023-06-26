package main

import (
    "leoferaderonugraha/go-backend-boilerplate/migrations"

    "flag"
)

func main() {
    face := flag.String("face", "up", "Specify the migration type (up/down)")

    migrations.Run(*face)
}
