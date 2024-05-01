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

func getProduct(id uint64) (Product, error) {
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

func createProduct(product Product) error {
	db := db.InitDbMysql()
	defer db.Close()

	_, err := db.NamedExec(`
		INSERT INTO product (name, price, description, created_at, updated_at)
		VALUES (:name, :price, :description, NOW(), NOW())
	`, product)

	if err != nil {
		logger.NewLogger().Error(err.Error())
		return errors.New("erro ao criar produto")
	}

	return nil
}

func updateProduct(product Product) error {
	db := db.InitDbMysql()
	defer db.Close()

	_, err := getProduct(product.ID)
	if err != nil {
		return err
	}

	_, err = db.NamedExec(`
		UPDATE product
		SET name = :name, price = :price, description = :description, updated_at = NOW()
		WHERE id = :id
	`, product)

	if err != nil {
		logger.NewLogger().Error(err.Error())
		return errors.New("erro ao atualizar produto")
	}

	return nil
}

func deleteProduct(id uint64) error {
	db := db.InitDbMysql()
	defer db.Close()

	_, err := getProduct(id)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		UPDATE product
		SET deleted_at = NOW()
		WHERE id = ?
	`, id)

	if err != nil {
		logger.NewLogger().Error(err.Error())
		return errors.New("erro ao deletar produto")
	}

	return nil
}
