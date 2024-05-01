package products

import (
	"github.com/natajonasdacoliveira/eulabs-challenge/db"
	"github.com/natajonasdacoliveira/eulabs-challenge/logger"
)

func getProducts(productName string) ([]Product, error) {
	db := db.InitDbMysql()
	defer db.Close()

	products := []Product{}

	err := db.Select(&products, `
		SELECT * FROM product
		WHERE deleted_at IS NULL
		AND name LIKE ?
	`, "%"+productName+"%")
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
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
		return Product{}, err
	}
	return product, nil
}
