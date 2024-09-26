package song_storage

import (
	"database/sql"
	"fmt"
	storage_models "github.com/Olegsuus/SongApi/internal/storage/models"
)

func (s *SongStorage) GetText(id int) (*storage_models.Song, error) {
	const op = "song_storage.GetText"

	query := `SELECT id, "group", song, release_date, text, link, created_at, updated_at FROM songs WHERE id = $1`

	var song storage_models.Song
	err := s.DB.QueryRow(query, id).Scan(
		&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Song with id %d not found", id)
		}
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	return &song, nil
}
