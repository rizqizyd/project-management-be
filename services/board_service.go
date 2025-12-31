package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rizqizyd/project-management-be/models"
	"github.com/rizqizyd/project-management-be/repositories"
)

type BoardService interface {
	Create(board *models.Board) error
	Update(board *models.Board) error
	GetByPublicID(publicID string) (*models.Board, error)
	AddMembers(boardPublicID string, userPublicIDs []string) error
	RemoveMembers(boardPublicID string, userPublicIDs []string) error
	GetAllByUserPaginate(userID, filter, sort string, limit, offset int) ([]models.Board, int64, error)
}

type boardService struct {
	boardRepo       repositories.BoardRepository
	userRepo        repositories.UserRepository
	boardMemberRepo repositories.BoardMemberRepository
}

func NewBoardService(boardRepo repositories.BoardRepository, userRepo repositories.UserRepository, boardMemberRepo repositories.BoardMemberRepository) BoardService {
	return &boardService{boardRepo, userRepo, boardMemberRepo}
}

func (s *boardService) Create(board *models.Board) error {
	user, err := s.userRepo.FindByPublicID(board.OwnerPublicID.String())
	if err != nil {
		return errors.New("owner not found")
	}

	board.PublicID = uuid.New()
	board.OwnerID = user.InternalID
	return s.boardRepo.Create(board)
}

func (s *boardService) Update(board *models.Board) error {
	return s.boardRepo.Update(board)
}

func (s *boardService) GetByPublicID(publicID string) (*models.Board, error) {
	return s.boardRepo.FindByPublicID(publicID)
}

func (s *boardService) AddMembers(boardPublicID string, userPublicIDs []string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDs {
		user, err := s.userRepo.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("user not found: " + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}

	// Check for existing members to avoid duplicate entries
	existingMembers, err := s.boardMemberRepo.GetMembers(string(board.PublicID.String()))
	if err != nil {
		return err
	}

	// Fast check existing member IDs using a map
	existingMemberIDs := make(map[uint]bool)
	for _, member := range existingMembers {
		existingMemberIDs[uint(member.InternalID)] = true // existingMemberIDs[member.InternalID] = true
	}

	var newMemberInternalIDs []uint
	for _, userInternalID := range userInternalIDs {
		if !existingMemberIDs[userInternalID] {
			newMemberInternalIDs = append(newMemberInternalIDs, userInternalID)
		}
	}
	if len(newMemberInternalIDs) == 0 {
		return nil // No new members to add
	}

	return s.boardRepo.AddMember(uint(board.InternalID), newMemberInternalIDs)
}

func (s *boardService) RemoveMembers(boardPublicID string, userPublicIDs []string) error {
	board, err := s.boardRepo.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	// Validate users and get their internal IDs
	var userInternalIDs []uint
	for _, userPublicID := range userPublicIDs {
		user, err := s.userRepo.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("user not found: " + userPublicID)
		}
		userInternalIDs = append(userInternalIDs, uint(user.InternalID))
	}

	// Check existing members to avoid removing non-members
	existingMembers, err := s.boardMemberRepo.GetMembers(string(board.PublicID.String()))
	if err != nil {
		return err
	}

	// Fast check existing member IDs using a map and only keep those that are members
	existingMemberIDs := make(map[uint]bool)
	for _, member := range existingMembers {
		existingMemberIDs[uint(member.InternalID)] = true // existingMemberIDs[member.InternalID] = true
	}

	var membersToRemove []uint
	for _, userInternalID := range userInternalIDs {
		if existingMemberIDs[userInternalID] {
			membersToRemove = append(membersToRemove, userInternalID)
		}
	}

	return s.boardRepo.RemoveMembers(uint(board.InternalID), membersToRemove)
}

func (s *boardService) GetAllByUserPaginate(userID, filter, sort string, limit, offset int) ([]models.Board, int64, error) {
	return s.boardRepo.FindAllByUserPaginate(userID, filter, sort, limit, offset)
}
