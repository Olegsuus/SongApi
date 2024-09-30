package handlers

import song_handlers "github.com/Olegsuus/SongApi/internal/handlers/song"

type Handler struct {
	SongHandler *song_handlers.SongHandlers
}

func NewHandler(songHandlers *song_handlers.SongHandlers) *Handler {
	return &Handler{
		SongHandler: songHandlers,
	}
}
