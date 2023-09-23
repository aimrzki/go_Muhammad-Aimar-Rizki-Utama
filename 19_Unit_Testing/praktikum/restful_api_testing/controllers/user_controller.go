package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"restful_api_testing/authentication"
	"restful_api_testing/config"
	"restful_api_testing/models"
	"strconv"
)

func GetUsersController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token is required")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token is required")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid user ID in token claims")
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to access this user data")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := authentication.CreateJWTToken(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create JWT token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
		"token":   token,
	})
}

func LoginUserController(c echo.Context) error {
	request := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	if user.Password != request.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	token, err := authentication.CreateJWTToken(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create JWT token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"user_id": user.ID,
		"token":   token,
	})
}

func DeleteUserController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token is required")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid user ID in token claims")
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to delete this user account")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func UpdateUserController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token is required")
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid user ID in token claims")
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to update this user account")
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := config.DB.Save(&existingUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    existingUser,
	})
}
