package products

import (
	"github.com/natajonasdacoliveira/eulabs-challenge/db"
	"github.com/natajonasdacoliveira/eulabs-challenge/logger"
)

func getProducts() ([]Product, error) {
	db := db.InitDbMysql()
	defer db.Close()

	products := []Product{}

	err := db.Select(&products, `
		SELECT * FROM product
		WHERE deleted_at IS NULL
	`)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil, err
	}
	return products, nil
}
