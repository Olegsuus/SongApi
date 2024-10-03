package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/handlers/response"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

// Add     		 godoc
// @Summary      Add a new song
// @Description  Добавить новую песню по DTO. Макс. длина песни и группы 255, допускается только Латиница
// @Tags         Song
// @Accept       json
// @Produce      json
// @Param        song  body      models.AddSongDTO  true  "Song creation data"
// @Success      201   {object}  response.AddSongResponse "Successfully added new song"
// @Failure      400   "Invalid data for adding the song"
// @Failure      500   "Failed to add the song"
// @Router       /song [post]
func (h *SongHandlers) Add(c echo.Context) error {
	const op = "song_handlers.add"

	h.l.With(slog.String("op", op))

	var dto models.AddSongDTO
	if err := c.Bind(&dto); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Ошибочные данные для добавления")
	}

	if err := h.v.Struct(dto); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Ошибка валидации, сверьтесь со swagger")
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
