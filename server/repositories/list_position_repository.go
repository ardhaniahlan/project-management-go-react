package repositories

import (
	"errors"
	"project-management-be/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListPositionRepository interface {
	GetByBoard(boardPublicID string) (*models.ListPosition, error)
	CreateOrUpdate(boardPublicID string, listOrder []uuid.UUID) error
	GetListOrder(boardPublicID string) ([]uuid.UUID, error)
	UpdateListOrder(position *models.ListPosition) error
}

type listPositionRepository struct {
	db *gorm.DB
}

func NewListPositionRepository(db *gorm.DB) ListPositionRepository {
	return &listPositionRepository{db: db}
}

func (r *listPositionRepository) GetByBoard(boardPublicID string) (*models.ListPosition, error) {
	var position models.ListPosition
	err := r.db.Where("board_internal_id = (SELECT internal_id FROM boards WHERE public_id = ?)", boardPublicID).
		First(&position).Error
	return &position, err
}

func (r *listPositionRepository) CreateOrUpdate(boardPublicID string, listOrder []uuid.UUID) error {
	return r.db.Exec(`
		INSERT INTO list_position (board_internal_id, list_order)
		SELECT internal_id, ? FROM boards WHERE public_id = ?
		ON CONFLICT (board_internal_id)
		DO UPDATE SET list_order = EXCLUDED.list_order;
	`, listOrder, boardPublicID).Error
}

func (r *listPositionRepository) GetListOrder(boardPublicID string) ([]uuid.UUID, error) {
	position, err := r.GetByBoard(boardPublicID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []uuid.UUID{}, nil
	}
	if err != nil {
		return nil, err
	}
	return position.ListOrder, nil
}

func (r *listPositionRepository) UpdateListOrder(position *models.ListPosition) error {
	return r.db.Model(&models.ListPosition{}).Where("internal_id = ?", position.InternalID).Update("list_order", position.ListOrder).Error
}
