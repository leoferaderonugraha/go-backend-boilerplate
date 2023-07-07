package handlers

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/services"
    e "leoferaderonugraha/go-backend-boilerplate/pkg/errors"

	"github.com/gofiber/fiber/v2"
    "errors"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userRegistrationService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userRegistrationService,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var request services.UserRegistrationRequest
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.userService.Register(request.Name, request.Email, request.Password)
	if err != nil {
        if errors.Is(err, e.USER_ALREADY_EXISTS) {
            return fiber.NewError(fiber.StatusConflict, err.Error())
        } else {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
	}

	response := services.UserRegistrationResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
    return nil
}
