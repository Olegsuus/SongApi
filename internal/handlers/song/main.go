package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/models"
	"log/slog"
)

type SongHandlers struct {
	l       *slog.Logger
	Service SongHandlersProvider
}

type SongHandlersProvider interface {
	Add(group, song string) (*models.Song, error)
	Update(song *models.Song) error
	GetText(id, page, size int) (*models.SongText, error)
	GetMany(getManyS models.GetManySong, limit, offset int, sortFields []string, isAscending bool) ([]*models.Song, error)
	Remove(id int) error
}

func NewSongHandlers(service SongHandlersProvider, l *slog.Logger) *SongHandlers {
	return &SongHandlers{
		Service: service,
		l:       l,
	}
}
