package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/rizqizyd/project-management-be/config"
	"github.com/rizqizyd/project-management-be/controllers"
	"github.com/rizqizyd/project-management-be/utils"
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
	app.Post("/v1/auth/login", uc.Login)

	// JWT Protected Routes
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Unauthorized(c, "Error unauthorized", err.Error())
		},
	}))

	userGroup := api.Group("/users")
	userGroup.Get("/:id", uc.GetUser)
}
