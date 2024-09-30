package handlers

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterRouters(e *echo.Echo) {
	e.POST("/song", h.SongHandler.Add)
	e.GET("/song/:id", h.SongHandler.GetSongText)
	e.GET("/songs", h.SongHandler.GetMany)
	e.PATCH("/songs/:id", h.SongHandler.Update)
	e.DELETE("/song/:id", h.SongHandler.Remove)
}
