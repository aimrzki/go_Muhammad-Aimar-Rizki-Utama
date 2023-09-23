package main

import (
	"ORM_and_Code_Structure_Full/config"
	"ORM_and_Code_Structure_Full/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	e := echo.New()
	routes.SetUserRoutes(e)
	routes.SetBookRoutes(e)
	routes.SetBlogRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
