package repositories

import (
	"strings"

	"github.com/rizqizyd/project-management-be/config"
	"github.com/rizqizyd/project-management-be/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByPublicID(publicID string) (*models.User, error)
	FindAllPagination(filter, sort string, limit, offset int) ([]models.User, int64, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByPublicID(publicID string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("public_id = ?", publicID).First(&user).Error
	return &user, err
}

func (r *userRepository) FindAllPagination(filter, sort string, limit, offset int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	db := config.DB.Model(&models.User{})

	// Filtering
	if filter != "" {
		filterPattern := "%" + filter + "%"
		db = db.Where("name ILIKE ? OR email ILIKE ?", filterPattern, filterPattern)
	}

	// Count total records
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting
	// if sort != "" {
	// 	db = db.Order(sort)
	// } else {
	// 	db = db.Order("created_at DESC")
	// }
	if sort != "" {
		switch sort {
		case "-id":
			sort = "-internal_id"
		case "id":
			sort = "internal_id"
		}

		if strings.HasPrefix(sort, "-") {
			sort = strings.TrimPrefix(sort, "-") + " DESC"
		} else {
			sort = sort + " ASC"
		}

		db = db.Order(sort)
	}

	err := db.Limit(limit).Offset(offset).Find(&users).Error
	return users, total, err
}

// Update updates an existing user record identified by public_id.
// It updates only the provided fields and avoids using the primary key directly.
func (r *userRepository) Update(user *models.User) error {
	// Model specifies the target table without binding to a specific record
	return config.DB.Model(&models.User{}).
		Where("public_id = ?", user.PublicID).
		Updates(map[string]interface{}{
			"name": user.Name,
		}).Error
}

func (r *userRepository) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
