package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/rizqizyd/project-management-be/config"
)

func CorsMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		// AllowOrigins:     config.AppConfig.CORSOrigin,
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length,Content-Type",
		MaxAge:           3000, // Cache preflight request for 50 minutes
	})
}
