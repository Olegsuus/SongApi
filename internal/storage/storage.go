package storage

import (
	"database/sql"
	storageModels "github.com/Olegsuus/SongApi/internal/storage/song"
)

type Storage struct {
	SongStorage *storageModels.SongStorage
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		SongStorage: storageModels.NewSongStorage(db),
	}
}
