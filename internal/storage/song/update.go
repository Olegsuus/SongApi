package song_storage

import (
	"fmt"
	storage_models "github.com/Olegsuus/SongApi/internal/storage/models"
)

func (s *SongStorage) Update(song *storage_models.Song) error {
	const op = "song_storage.update"

	query := `
		UPDATE songs 
		SET "group" = $1, song = $2, release_date = $3, text = $4, link = $5, updated_at = NOW()
		WHERE id = $6`

	result, err := s.DB.Exec(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, song.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to retrieve affected rows: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("song with id %d not found", song.ID)
	}

	return nil
}
