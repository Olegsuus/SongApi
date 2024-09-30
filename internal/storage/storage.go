package storage

import (
	"database/sql"
	song_storage "github.com/Olegsuus/SongApi/internal/storage/song"
)

type Storage struct {
	SongStorage *song_storage.SongStorage
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		SongStorage: song_storage.NewSongStorage(db),
	}
}
