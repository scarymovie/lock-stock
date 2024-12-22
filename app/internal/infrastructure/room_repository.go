package infrastructure

import (
	"errors"
	"lock-stock/internal/domain"
	"sync"
)

type InMemoryRoomRepository struct {
	mu    sync.RWMutex
	rooms map[string]*domain.Room
}

func NewInMemoryRoomRepository() domain.RoomRepository {
	return &InMemoryRoomRepository{
		rooms: make(map[string]*domain.Room),
	}
}

func (repo *InMemoryRoomRepository) FindByID(roomID string) (*domain.Room, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	room, exists := repo.rooms[roomID]
	if !exists {
		return nil, errors.New("room not found")
	}
	return room, nil
}

func (repo *InMemoryRoomRepository) Save(room *domain.Room) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.rooms[room.ID] = room
	return nil
}

func (repo *InMemoryRoomRepository) Delete(roomID string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.rooms[roomID]; !exists {
		return errors.New("room not found")
	}

	delete(repo.rooms, roomID)
	return nil
}
