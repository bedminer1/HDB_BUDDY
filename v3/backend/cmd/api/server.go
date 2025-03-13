package main

import (
	"github.com/labstack/echo/v4"
)

const version = "3.0.0"

func main() {
	e := echo.New()
	h := initHandler()

	e.GET("/healthcheck", h.handleHealthCheck)

	e.Logger.Fatal(e.Start(":4000"))
}