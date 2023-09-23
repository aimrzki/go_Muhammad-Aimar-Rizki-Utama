package routes

import (
	"ORM_and_Code_Structure_Full/controllers"
	"github.com/labstack/echo"
)

func SetBookRoutes(e *echo.Echo) {
	e.GET("/books", controllers.GetAllBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
}
