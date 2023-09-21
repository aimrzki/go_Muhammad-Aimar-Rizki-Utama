package controller

import (
	"clean_and_hexagonal_architecture/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (uc *UserController) CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	err := uc.db.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}
