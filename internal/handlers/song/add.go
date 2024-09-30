package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/handlers/response"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

func (h *SongHandlers) Add(c echo.Context) error {
	const op = "song_handlers.add"

	h.l.With(slog.String("op", op))

	var dto models.AddSongDTO
	if err := c.Bind(&dto); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Ошибочные данные для добавления")
	}

	song, err := h.Service.Add(dto.Group, dto.Song)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Ошибка при добавлении музыки")
	}

	h.l.Info("Successful add new song")

	resp := response.AddSongResponse{
		ID: song.ID,
	}

	return c.JSON(http.StatusCreated, resp)
}
