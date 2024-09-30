package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *SongHandlers) GetSongText(c echo.Context) error {
	const op = "song_handlers.GetSongText"

	h.l.With(slog.String("op", op))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Invalid song ID")
	}

	var getSong models.GetSongText
	if err := c.Bind(&getSong); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Invalid request parameters")
	}

	if getSong.Page < 1 {
		getSong.Page = 1
	}
	if getSong.Size < 1 {
		getSong.Size = 1
	}

	text, err := h.Service.GetText(id, getSong.Page, getSong.Size)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Failed to retrieve song lyrics")
	}

	return c.JSON(http.StatusOK, text)
}
