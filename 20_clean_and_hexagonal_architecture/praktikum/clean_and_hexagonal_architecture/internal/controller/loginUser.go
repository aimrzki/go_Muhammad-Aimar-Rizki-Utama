package controller

import (
	"clean_and_hexagonal_architecture/internal/model"
	"clean_and_hexagonal_architecture/middleware"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func (uc *UserController) LoginUser(c echo.Context) error {
	credentials := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user := model.User{}
	err := uc.db.Where("email = ?", credentials.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	if user.Password != credentials.Password {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(&user, uc.secretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Login successful", "token": token})
}
