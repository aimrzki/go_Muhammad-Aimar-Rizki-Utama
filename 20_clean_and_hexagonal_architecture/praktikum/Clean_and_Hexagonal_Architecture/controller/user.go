package controller

import (
	"Clean_and_Hexagonal_Architecture/dto"
	"Clean_and_Hexagonal_Architecture/helper"
	"Clean_and_Hexagonal_Architecture/model"
	"Clean_and_Hexagonal_Architecture/repository"
	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	userRepo repository.UserRepository
}

func NewUserController(uRepo repository.UserRepository) *userController {
	return &userController{
		uRepo,
	}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userRepo.Find()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err = u.userRepo.Create(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	token, err := helper.CreateToken(user.Email)
	if err != nil {
		return c.JSON(401, echo.Map{
			"error": err.Error(),
		})
	}
	uRes := dto.DTOUser{
		user.Email,
		user.Password,
		token,
	}
	return c.JSON(200, echo.Map{
		"data": uRes,
	})

}
