package song_storage

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	storageModels "github.com/Olegsuus/SongApi/internal/storage/models"
	"strings"
)

func (s *SongStorage) GetMany(getManyS models.GetManySong, limit, offset int, sortFields []string, isAscending bool) ([]*storageModels.Song, error) {
	const op = "song_storage.get_many"

	validSortFields := map[string]string{
		"group":       "\"group\"",
		"song":        "song",
		"releaseDate": "release_date",
		"text":        "text",
		"link":        "link",
	}

	sortClauses := []string{}
	for _, field := range sortFields {
		if column, ok := validSortFields[field]; ok {
			sortClauses = append(sortClauses, column)
		}
	}

	sortOrder := "ASC"
	if !isAscending {
		sortOrder = "DESC"
	}

	sortClause := ""
	if len(sortClauses) > 0 {
		sortClause = fmt.Sprintf("ORDER BY %s %s", strings.Join(sortClauses, ", "), sortOrder)
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
        %s
        LIMIT $6 OFFSET $7`, sortClause)

	rows, err := s.DB.Query(query, getManyS.Group, getManyS.Song, getManyS.ReleaseDate, getManyS.Text, getManyS.Link, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var songs []*storageModels.Song
	for rows.Next() {
		var song storageModels.Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link, &song.CreatedAt, &song.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		songs = append(songs, &song)
	}

	return songs, nil
}
