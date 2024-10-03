package errors

import (
	"github.com/labstack/echo/v4"
)

type ReqError struct {
	Status int    `json:"status"`
	Text   string `json:"text"`
}

func ErrorsHandler(c echo.Context, err error, status int, text string) error {
	if err != nil {
		return c.JSON(status, ReqError{
			Status: status,
			Text:   text,
		})
	}
	return nil
}
