package routes

import (
	"github.com/gofiber/fiber/v2"

	"ainyx-go-backend-task/internal/handler"
)

func RegisterUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/users")

	api.Post("/", userHandler.CreateUser)
	api.Get("/:id", userHandler.GetUserByID)
}
