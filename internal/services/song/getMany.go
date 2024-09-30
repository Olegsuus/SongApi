package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	"log/slog"
)

func (s *SongService) GetMany(group, song, releaseDate, text, link string, limit, offset int, sortBy, sortOrder string) ([]*models.Song, error) {
	const op = "song_services.get_many"

	s.l.With(slog.String("op", op))

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

	songsStorage, err := s.srP.GetMany(group, song, releaseDate, text, link, limit, offset, sortBy, sortOrder)
	if err != nil {
		s.l.Error("Failed to get songs", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var songs []*models.Song
	for _, songStorage := range songsStorage {
		song, err := s.TranslatorToModels(songStorage)
		if err != nil {
			s.l.Error("Failed to translate song storage to model", "error", err)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		songs = append(songs, song)
	}

	s.l.Info("Successfully retrieved songs")

	return songs, nil
}
