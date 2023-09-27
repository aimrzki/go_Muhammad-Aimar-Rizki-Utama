package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"restful_api_testing/config"
	"restful_api_testing/routes"
)

func main() {
	config.InitDB()
	e := echo.New()
	e.GET("/", routes.ServeHTML)
	routes.SetUserRoutes(e)
	routes.SetBookRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
