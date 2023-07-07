package handlers

import (
    "leoferaderonugraha/go-backend-boilerplate/src/app/repositories"
    "leoferaderonugraha/go-backend-boilerplate/src/app/services"
    e "leoferaderonugraha/go-backend-boilerplate/pkg/errors"

	"github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "errors"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
    repo := repositories.NewUserRepository(db)
    svc := services.NewUserRegistrationService(*repo)

    return &UserHandler{
        service: svc,
    }
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var request services.UserRegistrationRequest
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.service.Register(request.Name, request.Email, request.Password)
	if err != nil {
        if errors.Is(err, e.USER_ALREADY_EXISTS) {
            return fiber.NewError(fiber.StatusConflict, err.Error())
        } else {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "name": user.Name,
        "email": user.Email,
    })
}

func (h *UserHandler) Details(c *fiber.Ctx) error {
    id := c.Params("id")
    if id == "" {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
    }

    user, err := h.service.Details(id)
    if err != nil {
        if errors.Is(err, e.USER_NOT_FOUND) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "message": err.Error(),
            })
        } else {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Something went wrong",
            })
        }
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "name": user.Name,
        "email": user.Email,
    })
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
    id := c.Params("id")

    if id == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Bad Request",
        })
    }

    data := make(map[string]any)
    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }


    user, err := h.service.Update(id, data)

    if err != nil {
        if errors.Is(err, e.USER_NOT_FOUND) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "message": err.Error(),
            })
        } else {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": "Something went wrong",
            })
        }
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "name": user.Name,
        "email": user.Email,
    })
}
