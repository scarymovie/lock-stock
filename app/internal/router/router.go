package router

import (
	"lock-stock/internal/handler"
	"lock-stock/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter теперь принимает обработчик через инъекцию.
func NewRouter(joinRoomHandler *handler.JoinRoomHandler) http.Handler {
	r := chi.NewRouter()

	// Регистрация маршрута с параметром.
	r.With(middleware.LoggingMiddleware).Post("/room/{id}", joinRoomHandler.JoinRoom())

	return r
}
