package routes

import (
	"github.com/labstack/echo"
	"praktikum/controllers"
)

func SetUserRoutes(e *echo.Echo) {
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
}
