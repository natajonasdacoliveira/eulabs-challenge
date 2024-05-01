package products

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/natajonasdacoliveira/eulabs-challenge/logger"
	"github.com/natajonasdacoliveira/eulabs-challenge/validate"
)

func GetProducts(ctx echo.Context) error {
	productName := ctx.QueryParam("name")

	products, err := getProducts(productName)
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
			CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	baseDTO := BaseDTO{
		Data: productsDTO,
	}

	return ctx.JSON(http.StatusOK, baseDTO)
}

func GetProduct(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 19)

	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"errors": err.Error()})
	}

	product, err := getProduct(id)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	productDTO := ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	baseDTO := BaseDTO{
		Data: productDTO,
	}

	return ctx.JSON(http.StatusOK, baseDTO)
}

func CreateProduct(ctx echo.Context) error {
	product := Product{}
	err := ctx.Bind(&product)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"errors": err.Error()})
	}

	_, err = govalidator.ValidateStruct(product)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, validate.ValidateErrorWrapper(err))
	}

	err = createProduct(product)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"message": "produto criado com sucesso"})
}

func UpdateProduct(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 19)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"errors": err.Error()})
	}
	product := Product{}
	err = ctx.Bind(&product)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"errors": err.Error()})
	}

	_, err = govalidator.ValidateStruct(product)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, validate.ValidateErrorWrapper(err))
	}

	product.ID = id
	err = updateProduct(product)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "produto atualizado com sucesso"})
}

func DeleteProduct(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 19)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{"errors": err.Error()})
	}

	err = deleteProduct(id)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "produto deletado com sucesso"})
}
