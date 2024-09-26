package main

import (
	"github.com/Olegsuus/SongApi/internal/config"
	"github.com/Olegsuus/SongApi/internal/database"
	"github.com/Olegsuus/SongApi/internal/migrations"
)

func main() {
	cfg := config.GetConfig()
	db := database.DataBase{}

	db.GetStorage(cfg)
	migrations.Migrations(cfg, db.DB)

}
