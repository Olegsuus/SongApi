package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	"log/slog"
)

func (s *SongService) GetMany(group, song, releaseDate, text, link string, limit, offset int) ([]*models.Song, error) {
	const op = "song_services.get_all"

	s.l.With(slog.String("op", op))

	songsStorage, err := s.srP.GetMany(group, song, releaseDate, text, link, limit, offset)
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

	s.l.Info("Successful get all songs")

	return songs, nil
}
