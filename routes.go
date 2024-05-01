package main

import (
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	productsRoutes(e)
}

func productsRoutes(e *echo.Echo) {
	p := e.Group("/products")

	p.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "Hello, World",
		})
	})

	p.GET("/:id", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "Hello, Worldwith id",
		})
	})
}
