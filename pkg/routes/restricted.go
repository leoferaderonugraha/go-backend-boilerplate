package routes

import (
    "leoferaderonugraha/go-backend-boilerplate/pkg/middlewares"
    "leoferaderonugraha/go-backend-boilerplate/src/app/handlers"

    "github.com/gofiber/fiber/v2"
)

func RegisterRestrictedRoutes(app *fiber.App, handler *handlers.RestrictedHandler) {
    app.Get("/restricted", middlewares.AuthMiddleware, handler.AuthTest)
}

