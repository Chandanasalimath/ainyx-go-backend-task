package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"ainyx-go-backend-task/internal/handler"
	"ainyx-go-backend-task/internal/repository"
	"ainyx-go-backend-task/internal/routes"
	"ainyx-go-backend-task/internal/service"
)

func main() {
	app := fiber.New()

	// dummy db (task purpose)
	var db repository.DBTX

	repo := repository.New(db)
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	routes.RegisterUserRoutes(app, userHandler)

	log.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
