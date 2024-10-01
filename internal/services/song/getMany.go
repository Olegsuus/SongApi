package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	"log/slog"
)

func (s *SongService) GetMany(getManyS models.GetManySong, limit, offset int, sortFields []string, isAscending bool) ([]*models.Song, error) {
	const op = "song_services.get_many"

	s.l.With(slog.String("op", op))

	songsStorage, err := s.srP.GetMany(getManyS, limit, offset, sortFields, isAscending)
	if err != nil {
		s.l.Error("Не удалось получить песни", "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var songs []*models.Song
	for _, songStorage := range songsStorage {
		song, err := s.TranslatorToModels(songStorage)
		if err != nil {
			s.l.Error("Не удалось преобразовать данные песни", "error", err)
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		songs = append(songs, song)
	}

	s.l.Info("Успешно получили список песен")

	return songs, nil
}
