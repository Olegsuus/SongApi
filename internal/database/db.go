package database

import (
	"database/sql"
	"fmt"
	"github.com/Olegsuus/SongApi/internal/config"
	_ "github.com/lib/pq"
	"log"
)

type DataBase struct {
	DB *sql.DB
}

func (db *DataBase) GetStorage(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	fmt.Println(dsn)
	var err error
	db.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	if err = db.DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

func (db *DataBase) Stop() error {
	if db.DB != nil {
		err := db.DB.Close()
		{
			if err != nil {
				log.Fatalf("Failed to closed database: %s", err)
				return err
			}
		}
	}
	return nil
}
