package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":3000")) // コンテナ側の開放ポートと一緒にすること
}
