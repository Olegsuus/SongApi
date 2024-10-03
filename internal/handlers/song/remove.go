package song_handlers

import (
	"database/sql"
	errors2 "errors"
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/handlers/response"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

// Remove        godoc
// @Summary      Remove a song
// @Description  Deletes a song by its ID
// @Tags         Song
// @Param        id   path      int  true  "Song ID"
// @Success      200  {object}  response.SuccessResponse "Successfully deleted the song"
// @Failure      400  "Invalid ID format"
// @Failure      404  "Song not found"
// @Failure      500  "Internal server error"
// @Router       /song/{id} [delete]
func (h *SongHandlers) Remove(c echo.Context) error {
	const op = "song_handlers.remove"

	h.l.With(slog.String("op", op))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Не верный id")
	}

	err = h.Service.Remove(id)
	if err != nil {
		if errors2.Is(err, sql.ErrNoRows) {
			return errors.ErrorsHandler(c, err, http.StatusNotFound, "Музыка не найдена")
		}
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Failed to delete song")
	}

	h.l.Info("Successful delete song")

	resp := response.SuccessResponse{
		Success: true,
	}

	return c.JSON(http.StatusOK, resp)
}
