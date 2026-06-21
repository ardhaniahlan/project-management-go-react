package repositories

import (
	"project-management-be/models"
	"gorm.io/gorm"
)

type BoardRepository interface {
	Create(board *models.Board) error
	Update(board *models.Board, publicID string) error
	FindByPublicID(publicID string) (*models.Board, error)
	AddMembers(members []models.BoardMember) error
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

func (r *boardRepository) Update(board *models.Board, publicID string) error {
	return r.db.Model(&models.Board{}).Where("public_id = ?", publicID).Updates(map[string]interface{}{
		"title":       board.Title,
		"description": board.Description,
		"due_date":    board.DueDate,
	}).Error
}

func (r *boardRepository) FindByPublicID(publicID string) (*models.Board, error) {
	var board models.Board
	err := r.db.Where("public_id = ?", publicID).First(&board).Error
	return &board, err
}

func (r *boardRepository) AddMembers(members []models.BoardMember) error {
	return r.db.Create(&members).Error
}