package repositories

import (
	"github.com/rizqizyd/project-management-be/config"
	"github.com/rizqizyd/project-management-be/models"
)

type BoardRepository interface {
	Create(board *models.Board) error
}

type boardRepository struct{}

func NewBoardRepository() BoardRepository {
	return &boardRepository{}
}

func (r *boardRepository) Create(board *models.Board) error {
	return config.DB.Create(board).Error
}
