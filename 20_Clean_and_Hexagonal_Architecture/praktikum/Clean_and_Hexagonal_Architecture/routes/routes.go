package routes

import (
	"Clean_and_Hexagonal_Architecture/constants"
	"Clean_and_Hexagonal_Architecture/controller"
	"Clean_and_Hexagonal_Architecture/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRoute(app *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	AppJWT := app.Group("/users")
	AppJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	AppJWT.GET("", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
}
