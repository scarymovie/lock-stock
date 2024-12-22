//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"lock-stock/internal/handler"
	"lock-stock/internal/infrastructure"
	"lock-stock/internal/router"
	"lock-stock/internal/usecase"
	"net/http"
)

// InitializeRouter связывает все зависимости и возвращает готовый http.Handler.
func InitializeRouter() (http.Handler, error) {
	wire.Build(
		// Инфраструктура
		infrastructure.NewInMemoryRoomRepository,

		// Юзкейсы
		usecase.NewRoomService,

		// Хэндлеры
		handler.NewJoinRoomHandler, // Это будет использовать RoomService

		// Роутер
		router.NewRouter, // Передаем инжектированный JoinRoomHandler
	)

	return nil, nil
}
