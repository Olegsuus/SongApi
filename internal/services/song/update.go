package song_services

import (
	"database/sql"
	"errors"
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
		if errors.Is(err, sql.ErrNoRows) {
			s.l.Error("Не найден объект с таким id")
			return err
		} else {
			s.l.Error("Не удалось обновить объект")
			return fmt.Errorf("%s: %w", err)
		}
	}

	s.l.Info("Successful updated song", "id", song.ID)

	return nil
}
