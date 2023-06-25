package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type RestrictedHandler struct {
}

func NewRestrictedHandler() *RestrictedHandler {
    return &RestrictedHandler{}
}

func (h *RestrictedHandler) AuthTest(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "message": "You are authorized to access this route",
    })
}
