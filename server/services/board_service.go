package services

import (
	"errors"
	"project-management-be/models"
	"project-management-be/repositories"

	"github.com/google/uuid"
)

type BoardService interface {
	Create(board *models.Board, userPublicID string) error
	Update(boardPublicID string, board *models.Board, userPublicID string) error
	FindByPublicID(publicID string) (*models.Board, error)
}

type boardService struct {
	boardRepo repositories.BoardRepository
	userRepo  repositories.UserRepository
}

func NewBoardService(bRepo repositories.BoardRepository, uRepo repositories.UserRepository) BoardService {
	return &boardService{boardRepo: bRepo, userRepo: uRepo}
}

func (s *boardService) Create(board *models.Board, userPublicID string) error {
	user, err := s.userRepo.FindByPublicID(userPublicID)
	if err != nil {
		return errors.New("user not found")
	}

	board.PublicID = uuid.New()
	board.OwnerInternalID = user.InternalID
	board.OwnerPublicID = user.PublicID

	board.Owner = models.User{}

	return s.boardRepo.Create(board)
}

func (s *boardService) Update(boardPublicID string, board *models.Board, userPublicID string) error {
	user, err := s.userRepo.FindByPublicID(userPublicID)
	if err != nil {
		return errors.New("unauthorized: user not found")
	}

	existingBoard, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board tidak ditemukan")
	}

	if existingBoard.OwnerInternalID != user.InternalID {
		return errors.New("akses ditolak: anda bukan pemilik board ini")
	}

	board.InternalID = existingBoard.InternalID
	board.PublicID = existingBoard.PublicID
	board.OwnerInternalID = existingBoard.OwnerInternalID
	board.OwnerPublicID = existingBoard.OwnerPublicID
	board.CreatedAt = existingBoard.CreatedAt

	return s.boardRepo.Update(board, boardPublicID)
}

func (s *boardService) FindByPublicID(publicID string) (*models.Board, error) {
	return s.boardRepo.FindByPublicID(publicID)
}
