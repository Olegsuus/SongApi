package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	storagModels "github.com/Olegsuus/SongApi/internal/storage/models"
	"log/slog"
	"time"
)

func (s *SongService) Add(group, song string) (*models.Song, error) {
	const op = "song_services.add"

	s.l.With(slog.String("op", op))

	externalData, err := s.fetchSongDetails(group, song)
	if err != nil {
		s.l.Error("Не удалось получить данные о песне через внешние API", "group", group, "song", song, "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storageSong := &storagModels.Song{
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
		s.l.Error("Не удалось добавить песню в базу данных", "ошибка", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storageSong.ID = id
	modelSong, err := s.TranslatorToModels(storageSong)
	if err != nil {
		s.l.Error("Не удалось преобразовать данные из базы в модель", "ошибка", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	s.l.Info("Успешно добавлена новая песня", "id", modelSong.ID)

	return modelSong, nil
}
