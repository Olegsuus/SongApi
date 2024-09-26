package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	storage_models "github.com/Olegsuus/SongApi/internal/storage/models"
	"log/slog"
	"time"
)

func (s *SongService) Add(group, song string) (*models.Song, error) {
	const op = "song_services.add"

	s.l.With(slog.String("op", op))

	externalData, err := s.fetchSongDetails(group, song)
	if err != nil {
		s.l.Error("Failed to fetch song details from external API", "group", group, "song", song, "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storageSong := &storage_models.Song{
		Group:       group,
		Song:        song,
		ReleaseDate: externalData.ReleaseDate,
		Text:        externalData.Text,
		Link:        externalData.Link,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := s.srP.Add(storageSong)
	if err != nil {
		s.l.Error("Failed to add song to database", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storageSong.ID = id
	modelSong, err := s.TranslatorToModels(storageSong)
	if err != nil {
		s.l.Error("Failed to translate storage model to client model", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("Successfully added new song", "id", modelSong.ID)
	return modelSong, nil
}
