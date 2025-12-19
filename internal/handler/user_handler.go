package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"ainyx-go-backend-task/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid dob format (YYYY-MM-DD)",
		})
	}

	user, err := h.service.CreateUser(req.Name, dob)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	user, age, err := h.service.GetUserByID(int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
		"age":  age,
	})
}
