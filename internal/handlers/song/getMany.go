package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *SongHandlers) GetMany(c echo.Context) error {
	var getManySong models.GetManySong
	if err := c.Bind(&getManySong); err != nil {
		return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Invalid request parameters")
	}

	if getManySong.Page < 1 {
		getManySong.Page = 1
	}

	offset := (getManySong.Page - 1) * getManySong.Size

	songs, err := h.Service.GetMany(
		getManySong.Group, getManySong.Song,
		getManySong.ReleaseDate, getManySong.Text,
		getManySong.Link, getManySong.Size, offset,
	)

	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Failed to retrieve songs")
	}

	response := models.GetManySongs{
		Songs: songs,
		Page:  getManySong.Page,
		Size:  getManySong.Size,
	}

	return c.JSON(http.StatusOK, response)
}
