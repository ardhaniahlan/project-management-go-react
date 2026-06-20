package services

import (
	"errors"
	"project-management-be/models"
	"project-management-be/repositories"

	"github.com/google/uuid"
)

type BoardService interface {
	Create(board *models.Board, userPublicID string) error
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