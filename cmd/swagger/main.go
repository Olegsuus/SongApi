// @title         Your API Title
// @version       1.0
// @description   This is a sample server for a song API.
// @host          localhost:4444
// @BasePath      /
package main

import (
	_ "github.com/Olegsuus/SongApi/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
