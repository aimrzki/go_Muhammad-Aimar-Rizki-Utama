package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"
	"restful_api_testing/config"
	"restful_api_testing/routes"
)

func main() {
	config.InitDB()
	e := echo.New()
	routes.SetUserRoutes(e)
	routes.SetBookRoutes(e)
	http.HandleFunc("/", routes.ServeHTML)
	e.Logger.Fatal(e.Start(":8080"))
}
