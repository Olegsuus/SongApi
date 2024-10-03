package handlers

import songHandlers "github.com/Olegsuus/SongApi/internal/handlers/song"

type Handler struct {
	SongHandler *songHandlers.SongHandlers
}

func NewHandler(songHandlers *songHandlers.SongHandlers) *Handler {
	return &Handler{
		SongHandler: songHandlers,
	}
}
