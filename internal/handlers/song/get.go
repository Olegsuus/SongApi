package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

// GetSongText     godoc
// @Summary      Get song lyrics
// @Description  Retrieves the text (lyrics) of a song by its ID with pagination support
// @Tags         Song
// @Accept       json
// @Produce      json
// @Param        id    path     int     true  "Song ID"
// @Param        page  query    int     false "Page number (default is 1)"
// @Param        size  query    int     false "Page size (default is 5)"
// @Success      200   {object} string  "Successfully retrieved the song lyrics"
// @Failure      400   "Invalid song ID or request parameters"
// @Failure      500   "Failed to retrieve song lyrics"
// @Router       /song/{id} [get]
func (h *SongHandlers) GetSongText(c echo.Context) error {
	const op = "song_handlers.GetSongText"

	h.l.With(slog.String("op", op))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Invalid song ID")
	}

	var getSong models.GetSongText
	if err = c.Bind(&getSong); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Invalid request parameters")
	}

	if getSong.Page < 1 {
		getSong.Page = 1
	}

	text, err := h.Service.GetText(id, getSong.Page, getSong.Size)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Песня с таким id не найдена")
	}

	return c.JSON(http.StatusOK, text)
}
