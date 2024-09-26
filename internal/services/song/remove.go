package song_services

import (
	"fmt"
	"log/slog"
)

func (s *SongService) Remove(id int) error {
	const op = "song_services.delete"

	s.l.With(slog.String("op", op))

	err := s.srP.Remove(id)
	if err != nil {
		s.l.Error("Failed to delete song", "id", id, "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("Successful deleted song", "id", id)
	return nil
}
