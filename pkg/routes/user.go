package routes

import (
    "github.com/gofiber/fiber/v2"
    "leoferaderonugraha/go-backend-boilerplate/src/app/handlers"
)

func RegisterTestRoutes(app *fiber.App, handler *handlers.RestrictedHandler) {
    app.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World")
    })
}

func RegisterUserRoutes(app *fiber.App, handler *handlers.UserHandler) {
    group := app.Group("/user")
    group.Post("/register", handler.Register)
    group.Get("/:id", handler.Details)
    group.Put("/:id", handler.Update)
}
