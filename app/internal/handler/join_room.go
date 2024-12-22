package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"lock-stock/internal/usecase"
	"net/http"
)

// JoinRoomRequest - структура для входящего запроса.
type JoinRoomRequest struct {
	PlayerName string `json:"player_name"`
}

// JoinRoomHandler - структура обработчика, которая хранит зависимость от RoomService.
type JoinRoomHandler struct {
	roomService usecase.RoomService
}

// NewJoinRoomHandler - конструктор для создания обработчика с инъекцией зависимости.
func NewJoinRoomHandler(roomService usecase.RoomService) *JoinRoomHandler {
	return &JoinRoomHandler{
		roomService: roomService,
	}
}

// JoinRoom - метод для обработки запроса.
func (h *JoinRoomHandler) JoinRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлечение roomID из URL
		roomID := chi.URLParam(r, "id")
		if roomID == "" {
			http.Error(w, "Room ID is required", http.StatusBadRequest)
			return
		}

		var req JoinRoomRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Вызов RoomService.
		err := h.roomService.JoinRoom(roomID, req.PlayerName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Player %s joined room %s", req.PlayerName, roomID)
	}
}
