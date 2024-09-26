package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	"log/slog"
)

func (s *SongService) Update(song *models.Song) error {
	const op = "song_services.update"

	s.l.With(slog.String("op", op))

	storageSong, err := s.TranslatorToStorage(song)
	if err != nil {
		s.l.Error("Failed to translate to storage model", "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.srP.Update(storageSong)
	if err != nil {
		s.l.Error("Failed to update song", "id", song.ID, "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("Successful updated song", "id", song.ID)
	return nil
}
