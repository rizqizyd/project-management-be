package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rizqizyd/project-management-be/config"
	"github.com/rizqizyd/project-management-be/controllers"
	"github.com/rizqizyd/project-management-be/database/seed"

	// "github.com/rizqizyd/project-management-be/middleware"
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

	// Middleware
	// app.Use(middleware.CorsMiddleware())

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:5173", // Your frontend URLs
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// User
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	// List
	listRepo := repositories.NewListRepository()
	listPosRepo := repositories.NewListPositionRepository()
	listService := services.NewListService(listRepo, boardRepo, listPosRepo)
	listController := controllers.NewListController(listService)

	routes.Setup(app, userController, boardController, listController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port:", port)

	log.Fatal(app.Listen(":" + port))
}
