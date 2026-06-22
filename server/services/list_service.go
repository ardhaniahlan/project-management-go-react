package services

import (
	"errors"
	"fmt"
	"project-management-be/models"
	"project-management-be/models/types"
	"project-management-be/repositories"
	"project-management-be/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListService interface {
	GetByBoardID(boardPublicID string) (*ListWithOrder, error)
	GetByID(internalID uint) (*models.List, error)
	GetByPublicID(publicID string) (*models.List, error)
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(id uint) error
	UpdatePosition(boardPublicID string, positions []uuid.UUID) error
}

type ListWithOrder struct {
	Position []uuid.UUID
	Lists    []models.List
}

type listService struct {
	db          *gorm.DB
	listRepo    repositories.ListRepository
	boardRepo   repositories.BoardRepository
	listPosRepo repositories.ListPositionRepository
}

func NewListService(db *gorm.DB, listRepo repositories.ListRepository, boardRepo repositories.BoardRepository, listPosRepo repositories.ListPositionRepository) ListService {
	return &listService{db: db, listRepo: listRepo, boardRepo: boardRepo, listPosRepo: listPosRepo}
}

func (s *listService) GetByBoardID(boardPublicID string) (*ListWithOrder, error) {
	_, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return nil, errors.New("board not found")
	}

	position, err := s.listPosRepo.GetListOrder(boardPublicID)
	if err != nil {
		return nil, errors.New("Failed to get list order: " + err.Error())
	}

	lists, err := s.listRepo.FindByBoardID(boardPublicID)
	if err != nil {
		return nil, errors.New("Failed to get list: " + err.Error())
	}

	orderedList := utils.SortingListsByPosition(lists, position)

	return &ListWithOrder{
		Position: position,
		Lists:    orderedList,
	}, nil
}

func (s *listService) GetByID(internalID uint) (*models.List, error) {
	return s.listRepo.FindByID(internalID)
}

func (s *listService) GetByPublicID(publicID string) (*models.List, error) {
	return s.listRepo.FindByPublicID(publicID)
}

func (s *listService) Create(list *models.List) error {

	board, err := s.boardRepo.FindByPublicID(list.BoardPublicID.String())
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return errors.New("board not found")
		}
		return fmt.Errorf("failed to get board: %v", err)
	}

	list.BoardInternalID = int64(board.InternalID)
	if list.PublicID == uuid.Nil {
		list.PublicID = uuid.New()
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(list).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create list: %v", err)
	}

	var position models.ListPosition
	if err := tx.Where("board_internal_id = ?", board.InternalID).First(&position).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			position = models.ListPosition{
				PublicID:        uuid.New(),
				BoardInternalID: int64(board.InternalID),
				ListOrder:       types.UUIDArray{list.PublicID},
			}
			if err := tx.Create(&position).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to create list position: %v", err)
			}
		}
	} else {
		position.ListOrder = append(position.ListOrder, list.PublicID)
		if err := tx.Model(&position).Update("list_order", position.ListOrder).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update list position: %v", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (s *listService) Update(list *models.List) error {
	return s.listRepo.Update(list)
}

func (s *listService) Delete(id uint) error {
	return s.listRepo.Delete(id)
}

func (s *listService) UpdatePosition(boardPublicID string, positions []uuid.UUID) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	position, err := s.listPosRepo.GetByBoard(board.PublicID.String())
	if err != nil {
		return errors.New("list position not found")
	}

	position.ListOrder = positions
	return s.listPosRepo.UpdateListOrder(position)
}
