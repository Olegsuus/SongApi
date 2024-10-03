package song_services

import (
	"github.com/Olegsuus/SongApi/internal/models"
	storageModels "github.com/Olegsuus/SongApi/internal/storage/models"
	"log/slog"
)

type SongService struct {
	l   *slog.Logger
	srP ServiceProvider
}

type ServiceProvider interface {
	Add(song *storageModels.Song) (int, error)
	GetText(id int) (*storageModels.Song, error)
	GetMany(getManyS models.GetManySong, limit, offset int, sortFields []string, isAscending bool) ([]*storageModels.Song, error)
	Update(song *storageModels.Song) error
	Remove(id int) error
}

func NewSongService(l *slog.Logger, srP ServiceProvider) *SongService {
	return &SongService{
		l:   l,
		srP: srP,
	}
}
