package domain

import "errors"

// Room - структура комнаты.
type Room struct {
	ID      string
	Players []string
	MaxSize int
}

// RoomRepository - интерфейс репозитория.
type RoomRepository interface {
	FindByID(roomID string) (*Room, error)
	Save(room *Room) error
}

// AddPlayer - добавление игрока в комнату.
func (r *Room) AddPlayer(playerName string) error {
	if len(r.Players) >= r.MaxSize {
		return errors.New("room is full")
	}

	for _, player := range r.Players {
		if player == playerName {
			return errors.New("player already in room")
		}
	}

	r.Players = append(r.Players, playerName)
	return nil
}
