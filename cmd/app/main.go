package main

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/models"
    "leoferaderonugraha/go-backend-boilerplate/pkg/config"
    "leoferaderonugraha/go-backend-boilerplate/pkg/database"
    "leoferaderonugraha/go-backend-boilerplate/pkg/routes"
    "leoferaderonugraha/go-backend-boilerplate/src/app/handlers"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "fmt"
)

func main() {
    cfg, err := config.GetConfig()

    if err != nil {
        panic(err)
    }

    fmt.Println("Hello, World!")
    fmt.Println(cfg.DatabaseURL)

    return

    conn := database.NewDb()
    conn.Connect(cfg)

    database.WithTx(cfg, func (tx *gorm.DB) error {
        return tx.AutoMigrate(&models.User{})
    })

    app := fiber.New()

    userHandler := handlers.NewUserHandler(conn.Db)

    restrictedHandler := handlers.NewRestrictedHandler()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World")
    })

    routes.RegisterUserRoutes(app, userHandler)
    routes.RegisterTestRoutes(app, nil)
    routes.RegisterRestrictedRoutes(app, restrictedHandler)

    app.Listen(":3000")
}
