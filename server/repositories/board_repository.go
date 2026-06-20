package repositories

import (
	"project-management-be/models"

	"gorm.io/gorm"
)

type BoardRepository interface {
	Create(board *models.Board) error
}

type boardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &boardRepository{db: db}
}

func (r *boardRepository) Create(board *models.Board) error {
	return r.db.Create(&board).Error
}