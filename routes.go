package main

import (
	"github.com/labstack/echo/v4"
	"github.com/natajonasdacoliveira/eulabs-challenge/modules/products"
)

func Routes(e *echo.Echo) {
	productsRoutes(e)
}

func productsRoutes(e *echo.Echo) {
	p := e.Group("/products")

	p.GET("", products.GetProducts)

	p.GET("/:id", products.GetProduct)
}
