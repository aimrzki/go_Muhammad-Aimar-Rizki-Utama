package controller

import (
	"clean_and_hexagonal_architecture/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (uc *UserController) GetAllUsers(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token is required")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return []byte(uc.secretKey), nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	users := make([]model.User, 0)
	err = uc.db.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": users,
	})
}
