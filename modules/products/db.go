package products

import (
	"errors"

	"github.com/natajonasdacoliveira/eulabs-challenge/db"
	"github.com/natajonasdacoliveira/eulabs-challenge/logger"
)

func getProducts(productName string) ([]Product, error) {
	db := db.InitDbMysql()
	defer db.Close()

	products := []Product{}
	query := `
		SELECT * FROM product
		WHERE deleted_at IS NULL
		AND name LIKE ?
	`

	err := db.Select(&products, query, "%"+productName+"%")

	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, errors.New("erro ao buscar produtos")
	}
	return products, nil
}

func getProduct(id string) (Product, error) {
	db := db.InitDbMysql()
	defer db.Close()

	product := Product{}

	err := db.Get(&product, `
		SELECT * FROM product
		WHERE id = ? AND deleted_at IS NULL
	`, id)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		if err.Error() == "sql: no rows in result set" {
			return Product{}, errors.New("produto não encontrado")
		}
		return Product{}, err
	}
	return product, nil
}
