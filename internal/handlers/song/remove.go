package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/handlers/response"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

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
		if err.Error() == "song with id "+strconv.Itoa(id)+" not found" {
			return errors.ErrorsHandler(c, err, http.StatusNotFound, "Музыка не найдена")
		}
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Failed to delete song")
	}

	h.l.Info("Successful delete song")

	resp := response.SuccessResponse{
		Success: true,
	}

	return c.JSON(http.StatusNoContent, resp)
}
