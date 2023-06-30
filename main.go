package main

import (
	"fmt"
	"indochat/database"
	"indochat/pkg/mysql"
	"indochat/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file 😢😢😢")
	}

	e := echo.New()
	mysql.DatabaseInit()
	database.RunMigration()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")

	fmt.Println("Running on Port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))

}
