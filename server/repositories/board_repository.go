package repositories

import (
	"project-management-be/models"
	"project-management-be/utils"
	"strings"

	"gorm.io/gorm"
)

type BoardRepository interface {
	Create(board *models.Board) error
	Update(board *models.Board, publicID string) error
	FindByPublicID(publicID string) (*models.Board, error)
	AddMembers(members []models.BoardMember) error
	RemoveMembers(boardInternalID uint, userIDs []uint) error
	FindAllByUserPaginate(userInternalID uint, filter, sort string, limit, offset int) ([]models.Board, int64, error)
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

func (r *boardRepository) RemoveMembers(boardInternalID uint, userIDs []uint) error {
	if len(userIDs) == 0 {
		return nil
	}
	return r.db.
		Where("board_internal_id = ? AND user_internal_id IN (?)", boardInternalID, userIDs).
		Delete(&models.BoardMember{}).Error
}

func (r *boardRepository) FindAllByUserPaginate(userInternalID uint, filter, sort string, limit, offset int) ([]models.Board, int64, error) {
	var boards []models.Board

	query := r.db.Model(&models.Board{})

	switch strings.ToLower(filter) {
	case "owner":
		query = query.Where("owner_internal_id = ?", userInternalID)

	case "participant":
		query = query.Where("internal_id IN (SELECT board_internal_id FROM board_members WHERE user_internal_id = ?) AND owner_internal_id != ?", userInternalID, userInternalID)

	default:
		query = query.Where("owner_internal_id = ? OR internal_id IN (SELECT board_internal_id FROM board_members WHERE user_internal_id = ?)", userInternalID, userInternalID)
	}

	safeSort := "created_at DESC"
	switch strings.ToLower(sort) {
	case "title_asc":
		safeSort = "title ASC"
	case "title_desc":
		safeSort = "title DESC"
	}
	query = query.Order(safeSort)

	count, err := utils.Paginate(query, limit, offset, "", &boards)
	if err != nil {
		return nil, 0, err
	}

	return boards, count, nil
}
