package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"praktikum/config"
	"praktikum/routes"
)

func main() {
	config.InitDB()
	e := echo.New()
	routes.SetUserRoutes(e)
	routes.SetBookRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
