package song_storage

import (
	"fmt"
	storage_models "github.com/Olegsuus/SongApi/internal/storage/models"
)

func (s *SongStorage) GetMany(group, song, releaseDate, text, link string, limit, offset int, sortBy, sortOrder string) ([]*storage_models.Song, error) {
	const op = "song_storage.get_many"

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	validSortFields := map[string]bool{
		"id":           true,
		"group":        true,
		"song":         true,
		"release_date": true,
		"created_at":   true,
		"updated_at":   true,
	}

	if !validSortFields[sortBy] {
		sortBy = "created_at"
	}

	query := fmt.Sprintf(`
        SELECT id, "group", song, release_date, text, link, created_at, updated_at 
        FROM songs 
        WHERE 
            ($1 = '' OR "group" = $1) AND
            ($2 = '' OR song = $2) AND
            ($3 = '' OR release_date = $3) AND
            ($4 = '' OR text ILIKE '%%' || $4 || '%%') AND
            ($5 = '' OR link = $5)
        ORDER BY %s %s
        LIMIT $6 OFFSET $7`, sortBy, sortOrder)

	rows, err := s.DB.Query(query, group, song, releaseDate, text, link, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var songs []*storage_models.Song
	for rows.Next() {
		var song storage_models.Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		songs = append(songs, &song)
	}

	return songs, nil
}
