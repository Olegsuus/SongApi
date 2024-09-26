package song_storage

import "database/sql"

type SongStorage struct {
	DB *sql.DB
}

func NewSongStorage(db *sql.DB) *SongStorage {
	return &SongStorage{DB: db}
}
