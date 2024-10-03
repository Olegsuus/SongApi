package main

import (
	"github.com/Olegsuus/SongApi/internal/app"
	"github.com/Olegsuus/SongApi/internal/config"
	"github.com/Olegsuus/SongApi/internal/database"
	"github.com/Olegsuus/SongApi/internal/migrations"
	"github.com/Olegsuus/SongApi/internal/storage"
	"log/slog"
	"os"
)

func main() {
	cfg := config.GetConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db := database.DataBase{}
	db.GetStorage(cfg)

	defer func() {
		if err := db.Stop(); err != nil {
			logger.Error("Failed to close database", "error", err)
		}
	}()

	migrations.Migrations(db.DB)

	store := storage.NewStorage(db.DB)

	App := app.NewApp(cfg, store, logger)

	if err := App.Start(store); err != nil {
		logger.Error("Failed to start server", "error", err)
	}

	defer func() {
		if err := App.Stop(); err != nil {
			logger.Error("Failed to stop the application", "error", err)
		}
	}()
}
