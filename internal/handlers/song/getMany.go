package song_handlers

import (
	"github.com/Olegsuus/SongApi/internal/handlers/errors"
	"github.com/Olegsuus/SongApi/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetMany       godoc
// @Summary      Get a paginated list of songs
// @Description  Retrieves a list of songs with optional filters, pagination, and sorting
// @Tags         Song
// @Accept       json
// @Produce      json
// @Param        page           query   int     false  "Page number (default is 1)"
// @Param        size           query   int     false  "Page size (default is 5)"
// @Param        isAscending    query   bool    false  "Sorting order (true for ascending, false for descending)"
// @Param        group    		query   string  false  "Filter by group name"
// @Param        song     		query   string  false  "Filter by song name"
// @Param        release_date 	query string false  "Filter by release date"
// @Param        text     		query   string  false  "Filter by text"
// @Param        link     		query   string  false  "Filter by link"
// @Param        group          query   bool    false  "Sort by group name"
// @Param        song           query   bool    false  "Sort by song name"
// @Param        releaseDate    query   bool    false  "Sort by release date"
// @Param        text           query   bool    false  "Sort by text"
// @Param        link           query   bool    false  "Sort by link"
// @Success      200  {object}  models.GetManySongs "Successfully retrieved the list of songs"
// @Failure      400  "Invalid query parameters"
// @Failure      500  "Failed to retrieve the list of songs"
// @Router       /songs [get]
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
