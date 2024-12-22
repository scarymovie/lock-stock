package usecase

import (
	"errors"
	"lock-stock/internal/domain"
)

// RoomService - интерфейс.
type RoomService interface {
	JoinRoom(roomID, playerName string) error
}

// roomService - реализация.
type roomService struct {
	roomRepo domain.RoomRepository
}

// NewRoomService - конструктор.
func NewRoomService(repo domain.RoomRepository) RoomService {
	return &roomService{roomRepo: repo}
}

func (s *roomService) JoinRoom(roomID, playerName string) error {
	room, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return errors.New("room not found")
	}

	if err := room.AddPlayer(playerName); err != nil {
		return err
	}

	return s.roomRepo.Save(room)
}
