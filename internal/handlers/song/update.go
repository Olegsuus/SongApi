package song_handlers

import (
	"database/sql"
	errors2 "errors"
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/handlers/response"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

// Update        godoc
// @Summary      Update a song
// @Description  Обновить запись о песне по DTO. Макс. длина группы и песни 255, допустима только Латиница.
// @Tags         Song
// @Accept       json
// @Produce      json
// @Param        id   path      int                     true  "Song ID"
// @Param        song body      models.UpdateSongDTO    true  "Song Update Payload"
// @Success      200  {object}  response.SuccessResponse "Successfully updated the song"
// @Failure      400  "Invalid input data or invalid ID format"
// @Failure      404  "Song not found"
// @Failure      500  "Internal server error"
// @Router       /song/{id} [patch]
func (h *SongHandlers) Update(c echo.Context) error {
	const op = "song_handlers.update"

	h.l.With(slog.String("op", op))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Не верный формат id")
	}

	var dto models.UpdateSongDTO
	if err = c.Bind(&dto); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Ошибочные данные для изменения")
	}

	if err = h.v.Struct(dto); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Ошибка валидации, сверьтесь со swagger")
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
		if errors2.Is(err, sql.ErrNoRows) {
			return errors.ErrorsHandler(c, err, http.StatusNotFound, "Song not found")
		}
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Failed to update song")
	}

	resp := response.SuccessResponse{
		Success: true,
	}

	return c.JSON(http.StatusOK, resp)
}
