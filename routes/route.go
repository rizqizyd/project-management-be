package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rizqizyd/project-management-be/controllers"
)

// App is the main Fiber framework instance in which all API route endpoints are registered.
// 'uc' is the UserController that handles user-related requests.
// This function sets up user route registration in the Fiber app.
func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Post("/v1/auth/register", uc.Register)
}
