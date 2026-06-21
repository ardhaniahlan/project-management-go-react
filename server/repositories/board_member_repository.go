package repositories

import (
	"project-management-be/models"

	"gorm.io/gorm"
)

type BoardMemberRepository interface {
	GetMember(boardInternalID uint) ([]uint, error)
}

type boardMemberRepository struct {
	db *gorm.DB
}

func NewBoardMemberRepository(db *gorm.DB) BoardMemberRepository {
	return &boardMemberRepository{db: db}
}

func (r *boardMemberRepository) GetMember(boardInternalID uint) ([]uint, error) {
	var existingIDs []uint
	err := r.db.Model(&models.BoardMember{}).
		Where("board_internal_id = ?", boardInternalID).
		Pluck("user_internal_id", &existingIDs).Error
	return existingIDs, err
}
