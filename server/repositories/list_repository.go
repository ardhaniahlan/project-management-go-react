package repositories

import (
	"project-management-be/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListRepository interface {
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(publicID uint) error
	UpdatePosition(boardPublicID string, position []string) error
	GetCardPosition(listPublicID string) ([]uuid.UUID, error)
	FindByBoardID(boardID string) ([]models.List, error)
	FindByPublicID(publicID string) (*models.List, error)
	FindByID(internalID uint) (*models.List, error)
}

type listRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) ListRepository {
	return &listRepository{db: db}
}

func (r *listRepository) Create(list *models.List) error {
	return r.db.Create(list).Error
}

func (r *listRepository) Update(list *models.List) error {
	return r.db.Model(&models.List{}).Where("public_id = ?", list.PublicID).Updates(map[string]interface{}{
		"title": list.Title,
	}).Error
}

func (r *listRepository) Delete(publicID uint) error {
	return r.db.Delete(&models.List{}, publicID).Error
}

func (r *listRepository) UpdatePosition(boardPublicID string, position []string) error {
	return r.db.Model(&models.ListPosition{}).
		Where("board_internal_id = (SELECT internal_id FROM boards WHERE public_id = ?)", boardPublicID).
		Update("list_order", position).Error
}

func (r *listRepository) GetCardPosition(listPublicID string) ([]uuid.UUID, error) {
	var positions models.CardPosition
	err := r.db.
	Joins("JOIN lists ON list.internal_id = card_position.list_internal_id").
	Where("list.public_id = ?", listPublicID).First(&positions).Error
	return positions.CardOrder, err
}

func (r *listRepository) FindByBoardID(boardID string) ([]models.List, error) {
	var lists []models.List
	err := r.db.Where("board_public_id = ?", boardID).Order("internal_id ASC").Find(&lists).Error
	return lists, err
}

func (r *listRepository) FindByPublicID(publicID string) (*models.List, error) {
	var list models.List
	err := r.db.Where("public_id = ?", publicID).First(&list).Error
	return &list, err
}

func (r *listRepository) FindByID(internalID uint) (*models.List, error) {
	var list models.List
	err := r.db.Where("internal_id = ?", internalID).First(&list).Error
	return &list, err
}


