package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/hello", hello)

	e.GET("/ping", ping)

	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world!! v2")
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
