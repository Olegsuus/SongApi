package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	response "github.com/Olegsuus/SongApi/internal/handlers/response"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *SongHandlers) Update(c echo.Context) error {
	const op = "song_handlers.update"

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Не верный формат id")
	}

	var dto models.UpdateSongDTO
	if err = c.Bind(&dto); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Ошибочные данные для изменения")
	}

	song := &models.Song{
		ID:          id,
		Group:       dto.Group,
		Song:        dto.Song,
		ReleaseDate: dto.ReleaseDate,
		Text:        dto.Text,
		Link:        dto.Link,
	}

	err = h.Service.Update(song)
	if err != nil {
		if err.Error() == "song with id "+strconv.Itoa(id)+" not found" {
			return errors.ErrorsHandler(c, err, http.StatusNotFound, "Song not found")
		}
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Failed to update song")
	}

	resp := response.SuccessResponse{
		Success: true,
	}

	return c.JSON(http.StatusOK, resp)
}
