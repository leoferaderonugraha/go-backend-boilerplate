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
)

func main() {
    cfg, err := config.LoadConfig("config.json")

    if err != nil {
        panic(err)
    }

    conn := database.New()
    conn.Connect(cfg)

    tx := conn.Db.Begin()
    if tx.AutoMigrate(&models.User{}) != nil {
        tx.Rollback()
        panic(err)
    }
    tx.Commit()

    app := fiber.New()

    userRepo := repositories.New(conn.Db)
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
