package main

import (
	// controller "controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// e.POST("user", controller.Create())
	// e.GET("user/:id", controller.UserShow())
	// e.PUT("user/:id", controller.Update())
	// e.DELETE("user/:id", controller.Delete())
	return e
}
