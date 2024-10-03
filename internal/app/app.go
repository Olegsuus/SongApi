package app

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/config"
	"github.com/Olegsuus/SongApi/internal/handlers/handlers"
	songHandler "github.com/Olegsuus/SongApi/internal/handlers/song"
	songService "github.com/Olegsuus/SongApi/internal/services/song"
	"github.com/Olegsuus/SongApi/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
)

type App struct {
	Config   *config.Config
	Handlers *handlers.Handler
	Echo     *echo.Echo
	Logger   *slog.Logger
}

func NewApp(cfg *config.Config, store *storage.Storage, logger *slog.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
	}
}

func (a *App) InitializeHandlers(store *storage.Storage) {
	songService := songService.NewSongService(a.Logger, store.SongStorage)

	songHandler := songHandler.NewSongHandlers(songService, a.Logger)

	a.Handlers = handlers.NewHandler(songHandler)
}

func (a *App) Start(store *storage.Storage) error {
	a.Echo = echo.New()

	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())

	a.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
	}))

	a.InitializeHandlers(store)

	a.Handlers.RegisterRouters(a.Echo)

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	a.Logger.Info("Starting server", "address", addr)

	return a.Echo.Start(addr)
}

func (a *App) Stop() error {
	return a.Echo.Shutdown(nil)
}
