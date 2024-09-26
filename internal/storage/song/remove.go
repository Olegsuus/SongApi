package song_storage

import "fmt"

func (s *SongStorage) Remove(id int) error {
	const op = "song_storage.delete"

	query := `DELETE FROM songs WHERE id = $1`

	result, err := s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to retrieve affected rows: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("song with id %d not found", id)
	}

	return nil
}
