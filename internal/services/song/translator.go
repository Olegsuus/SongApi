package song_services

import (
	"github.com/Olegsuus/SongApi/internal/models"
	storage_models "github.com/Olegsuus/SongApi/internal/storage/models"
)

func (s *SongService) TranslatorToModels(song *storage_models.Song) (*models.Song, error) {
	return &models.Song{
		ID:          song.ID,
		Group:       song.Group,
		Song:        song.Song,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
		CreatedAt:   song.CreatedAt,
		UpdatedAt:   song.UpdatedAt,
	}, nil
}

func (s *SongService) TranslatorToStorage(song *models.Song) (*storage_models.Song, error) {
	return &storage_models.Song{
		ID:          song.ID,
		Group:       song.Group,
		Song:        song.Song,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
		CreatedAt:   song.CreatedAt,
		UpdatedAt:   song.UpdatedAt,
	}, nil
}
