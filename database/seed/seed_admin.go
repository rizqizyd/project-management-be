package seed

import (
	"log"

	"github.com/rizqizyd/project-management-be/config"
	"github.com/rizqizyd/project-management-be/models"
	"github.com/rizqizyd/project-management-be/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Super admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Println("Admin user seeded")
	}
}
