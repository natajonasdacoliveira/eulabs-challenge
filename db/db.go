package db

import (
	_ "database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/natajonasdacoliveira/eulabs-challenge/logger"
)

// TODO: test logging data after
// TODO: maybe check the singleton instance aswell
func InitDbMysql() *sqlx.DB {
	err := godotenv.Load(".env")
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil
	}
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))
	db, err := sqlx.Connect("mysql", connectionString)

	if err != nil {
		logger.NewLogger().Error(err.Error())
		return nil
	}

	return db
}
