package routes

import (
    "github.com/gofiber/fiber/v2"
    "leoferaderonugraha/go-backend-boilerplate/src/app/handlers"
)

func TestRoutes(app *fiber.App) {
    app.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World")
    })
}

func RegisterUserRoutes(app *fiber.App, handler *handlers.UserHandler) {
    app.Post("/register", handler.Register)
}
