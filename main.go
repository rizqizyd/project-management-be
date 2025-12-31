package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rizqizyd/project-management-be/config"
	"github.com/rizqizyd/project-management-be/controllers"
	"github.com/rizqizyd/project-management-be/database/seed"
	"github.com/rizqizyd/project-management-be/repositories"
	"github.com/rizqizyd/project-management-be/routes"
	"github.com/rizqizyd/project-management-be/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()

	// Initialize Fiber app to set up routes and start the server
	// Fiber.New() creates a new instance (http) of the Fiber framework
	app := fiber.New()

	// User
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app, userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port:", port)

	log.Fatal(app.Listen(":" + port))
}
