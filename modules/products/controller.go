package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natajonasdacoliveira/eulabs-challenge/logger"
)

func GetProducts(ctx echo.Context) error {
	products, err := getProducts()
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	productsDTO := []ProductDTO{}

	for _, product := range products {
		productsDTO = append(productsDTO, ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	baseDTO := BaseDTO{
		Data: productsDTO,
	}

	return ctx.JSON(http.StatusOK, baseDTO)
}
