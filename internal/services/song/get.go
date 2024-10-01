package song_services

import (
	"fmt"
	"github.com/Olegsuus/SongApi/internal/models"
	"log/slog"
	"regexp"
	"strings"
)

var verseSplitter = regexp.MustCompile(`(\r?\n){2,}`)

func (s *SongService) GetText(id, page, size int) (*models.SongText, error) {
	const op = "song_services.GetText"

	s.l.With(slog.String("op", op))

	songStorage, err := s.srP.GetText(id)
	if err != nil {
		s.l.Error("Failed to get song", "id", id, "error", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	verses := verseSplitter.Split(songStorage.Text, -1)

	var cleanedVerses []string
	for _, verse := range verses {
		verse = strings.TrimSpace(verse)
		if verse != "" {
			cleanedVerses = append(cleanedVerses, verse)
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
			Page:   page,
			Size:   size,
		}, nil
	}

	endIndex := startIndex + size
	if endIndex > totalVerses {
		endIndex = totalVerses
	}

	s.l.Info("Successfully retrieved song text")

	return &models.SongText{
		ID:     songStorage.ID,
		Group:  songStorage.Group,
		Song:   songStorage.Song,
		Lyrics: cleanedVerses[startIndex:endIndex],
		Page:   page,
		Size:   size,
	}, nil
}
