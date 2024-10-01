package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *SongHandlers) GetMany(c echo.Context) error {
	var getManySong models.GetManySong

	pageParam := c.QueryParam("page")
	sizeParam := c.QueryParam("size")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(sizeParam)
	if err != nil || size < 1 {
		size = 5
	}

	offset := (page - 1) * size

	sortFields := []string{}
	validSortFields := []string{"group", "song", "releaseDate", "text", "link"}

	for _, field := range validSortFields {
		if c.QueryParam(field) == "true" {
			sortFields = append(sortFields, field)
		}
	}

	isAscending := true
	isAscendingParam := c.QueryParam("isAscending")
	if isAscendingParam != "" {
		isAscending, err = strconv.ParseBool(isAscendingParam)
		if err != nil {
			return errors.ErrorsHandler(c, err, http.StatusBadRequest, "Параметр isAscending должен быть true или false")
		}
	}

	getManySong.Group = c.QueryParam("groupFilter")
	getManySong.Song = c.QueryParam("songFilter")
	getManySong.ReleaseDate = c.QueryParam("releaseDateFilter")
	getManySong.Text = c.QueryParam("textFilter")
	getManySong.Link = c.QueryParam("linkFilter")

	songs, err := h.Service.GetMany(getManySong, size, offset, sortFields, isAscending)

	if err != nil {
		return errors.ErrorsHandler(c, err, http.StatusInternalServerError, "Не удалось получить список песен")
	}

	response := models.GetManySongs{
		Songs: songs,
		Page:  page,
		Size:  size,
	}

	return c.JSON(http.StatusOK, response)
}
