package song_services

import (
	"github.com/Olegsuus/SongApi/internal/models"
	storage_models "github.com/Olegsuus/SongApi/internal/storage/models"
	"log/slog"
)

type SongService struct {
	l   *slog.Logger
	srP ServiceProvider
}

type ServiceProvider interface {
	Add(song *storage_models.Song) (int, error)
	GetText(id int) (*storage_models.Song, error)
	GetMany(getManyS models.GetManySong, limit, offset int, sortFields []string, isAscending bool) ([]*storage_models.Song, error)
	Update(song *storage_models.Song) error
	Remove(id int) error
}

func NewSongService(l *slog.Logger, srP ServiceProvider) *SongService {
	return &SongService{
		l:   l,
		srP: srP,
	}
}
