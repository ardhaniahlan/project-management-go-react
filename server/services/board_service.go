package services

import (
	"errors"
	"project-management-be/models"
	"project-management-be/repositories"
	"time"

	"github.com/google/uuid"
)

type BoardService interface {
	Create(board *models.Board, userPublicID string) error
	Update(boardPublicID string, board *models.Board, userPublicID string) error
	FindByPublicID(publicID string) (*models.Board, error)
	AddMembers(boardPublicID string, userPublicIDs []string, actorPublicID string) error
}

type boardService struct {
	boardRepo       repositories.BoardRepository
	userRepo        repositories.UserRepository
	boardMemberRepo repositories.BoardMemberRepository
}

func NewBoardService(
	bRepo repositories.BoardRepository, 
	uRepo repositories.UserRepository, 
	bmRepo repositories.BoardMemberRepository,
	) BoardService {
	return &boardService{boardRepo: bRepo, userRepo: uRepo, boardMemberRepo: bmRepo}
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

func (s *boardService) AddMembers(boardPublicID string, userPublicIDs []string, actorPublicID string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board tidak ditemukan")
	}

	actor, err := s.userRepo.FindByPublicID(actorPublicID)
	if err != nil || board.OwnerInternalID != actor.InternalID {
		return errors.New("akses ditolak: anda bukan pemilik board ini")
	}

	usersToAdd, _ := s.userRepo.FindManyByPublicIDs(userPublicIDs)
	existingIDs, _ := s.boardMemberRepo.GetMember(uint(board.InternalID))

	existingMap := make(map[uint]bool)
	for _, id := range existingIDs {
		existingMap[id] = true
	}

	var newMembers []models.BoardMember
	now := time.Now()

	for _, user := range usersToAdd {
		if !existingMap[uint(user.InternalID)] {
			newMembers = append(newMembers, models.BoardMember{
				BoardInternalID: uint64(board.InternalID),
				UserInternalID:  uint64(user.InternalID),
				JoinedAt:        now,
			})
		}
	}

	if len(newMembers) == 0 {
		return errors.New("semua pengguna tersebut sudah menjadi anggota atau ID tidak valid")
	}

	return s.boardRepo.AddMembers(newMembers)
}
