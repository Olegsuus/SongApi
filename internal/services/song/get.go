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
		s.l.Error("Failed to get song", "id", id, "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	verses := strings.Split(songStorage.Text, "\n\n")

	totalVerses := len(verses)
	startIndex := (page - 1) * size
	if startIndex >= totalVerses {
		return &models.SongText{
			ID:     songStorage.ID,
			Group:  songStorage.Group,
			Song:   songStorage.Song,
			Lyrics: []string{},
			Page:   page,
			Size:   size,
			Total:  totalVerses,
		}, nil
	}

	endIndex := startIndex + size
	if endIndex > totalVerses {
		endIndex = totalVerses
	}

	s.l.Info("Successful get text for song")

	return &models.SongText{
		ID:     songStorage.ID,
		Group:  songStorage.Group,
		Song:   songStorage.Song,
		Lyrics: verses[startIndex:endIndex],
		Page:   page,
		Size:   size,
		Total:  totalVerses,
	}, nil
}
