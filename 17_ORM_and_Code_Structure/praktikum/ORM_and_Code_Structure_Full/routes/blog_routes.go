package routes

import (
	"ORM_and_Code_Structure_Full/controllers"
	"github.com/labstack/echo"
)

func SetBlogRoutes(e *echo.Echo) {
	e.POST("/blogs", controllers.CreateBlogController)
	e.GET("/blogs", controllers.GetAllBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
}
