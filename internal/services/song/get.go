package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	"log/slog"
	"strings"
)

func (s *SongService) GetText(id, page, size int) (*models.SongText, error) {
	const op = "song_services.GetText"

	s.l.With(slog.String("op", op))

	songStorage, err := s.srP.GetText(id)
	if err != nil {
		s.l.Error("Не удалось получить песню", "id", id, "error", err)
		fmt.Errorf("%s: %w", op, err)
		return nil, err
	}

	cleanedText := strings.ReplaceAll(songStorage.Text, "\r\n", "\n")

	lines := strings.Split(cleanedText, "\n")

	var cleanedVerses []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleanedVerses = append(cleanedVerses, line)
		}
	}

	totalVerses := len(cleanedVerses)
	if size <= 0 {
		size = totalVerses
	}

	startIndex := (page - 1) * size
	if startIndex >= totalVerses {
		return &models.SongText{
			ID:     songStorage.ID,
			Group:  songStorage.Group,
			Song:   songStorage.Song,
			Lyrics: []string{},
			Link:   songStorage.Link,
			Page:   page,
			Size:   size,
		}, nil
	}

	endIndex := startIndex + size
	if endIndex > totalVerses {
		endIndex = totalVerses
	}

	s.l.Info("Успешно получили текст песни")

	return &models.SongText{
		ID:     songStorage.ID,
		Group:  songStorage.Group,
		Song:   songStorage.Song,
		Lyrics: cleanedVerses[startIndex:endIndex],
		Link:   songStorage.Link,
		Page:   page,
		Size:   size,
	}, nil
}
