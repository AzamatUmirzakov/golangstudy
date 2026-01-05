package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	message string
}

func main() {
	e := echo.New()

	e.GET("/", getHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func getHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{message: "Hello World!"})
}
