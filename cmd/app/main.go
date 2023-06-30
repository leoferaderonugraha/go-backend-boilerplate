package main

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/models"
    "leoferaderonugraha/go-backend-boilerplate/pkg/config"
    "leoferaderonugraha/go-backend-boilerplate/pkg/database"
    "leoferaderonugraha/go-backend-boilerplate/pkg/routes"
    "leoferaderonugraha/go-backend-boilerplate/src/app/handlers"
    "leoferaderonugraha/go-backend-boilerplate/src/app/services"
    "leoferaderonugraha/go-backend-boilerplate/src/app/repositories"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func main() {
    cfg, err := config.LoadConfig("config.json")

    if err != nil {
        panic(err)
    }

    conn := database.NewDb()
    conn.Connect(cfg)

    database.WithTx(cfg, func (tx *gorm.DB) error {
        return tx.AutoMigrate(&models.User{})
    })

    app := fiber.New()

    userRepo := repositories.NewUserRepository(conn.Db)
    userService := services.NewUserRegistrationService(*userRepo)
    userHandler := handlers.NewUserHandler(userService)

    restrictedHandler := handlers.NewRestrictedHandler()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World")
    })

    routes.RegisterUserRoutes(app, userHandler)
    routes.TestRoutes(app)
    routes.RegisterRestrictedRoutes(app, restrictedHandler)

    app.Listen(":3000")
}
