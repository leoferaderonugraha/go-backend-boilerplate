package handlers

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/models"
    "leoferaderonugraha/go-backend-boilerplate/src/app/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRegistrationService *services.UserRegistrationService
}

func NewUserHandler(userRegistrationService *services.UserRegistrationService) *UserHandler {
	return &UserHandler{
		userRegistrationService: userRegistrationService,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var request models.UserRegistrationRequest
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.userRegistrationService.RegisterUser(request.Name, request.Email, request.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := models.UserRegistrationResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
    return nil
}
