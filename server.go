package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	InitializeAPI()
}

func InitializeAPI() {
	file, err := os.OpenFile("storage/echo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Error at opening log file: %v", err))
	}

	defer file.Close()

	API := echo.New()
	API.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:9000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	API.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: file,
	}))
	API.Use(middleware.Recover())

	Routes(API)

	API.Logger.Fatal(API.Start("127.0.0.1:3001"))
}
