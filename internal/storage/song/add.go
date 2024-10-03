package song_storage

import (
	"fmt"
	storageModels "github.com/Olegsuus/SongApi/internal/storage/models"
)

func (s *SongStorage) Add(song *storageModels.Song) (int, error) {
	const op = "song_storage.add"

	query := `
		INSERT INTO songs ("group", song, release_date, text, link, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id`

	var id int
	err := s.DB.QueryRow(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
