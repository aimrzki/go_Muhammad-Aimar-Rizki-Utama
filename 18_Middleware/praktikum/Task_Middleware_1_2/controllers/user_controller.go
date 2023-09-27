package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"praktikum/authentication"
	"praktikum/config"
	"praktikum/models"
	"strconv"
	"strings"
)

func GetUsersController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}

	tokenParts := strings.Split(authorization, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token harus dalam format 'Bearer <token>'")
	}

	token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses mendapatkan semua data user",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}

	tokenParts := strings.Split(authorization, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token harus dalam format 'Bearer <token>'")
	}

	token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Kesalahan dalam token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Token tidak sesuai dengan user Id")
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Id salah / tidak ditemukan")
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak memiliki akses untuk mengakses data ini")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses mendapatkan data user",
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
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal dalam pembuatan JWT token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil membuat akun",
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
		return echo.NewHTTPError(http.StatusUnauthorized, " email atau password tidak sesuai")
	}

	if user.Password != request.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "email atau password tidak sesuai")
	}

	token, err := authentication.CreateJWTToken(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal dalam pembuatan JWT token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login berhasil",
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
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}

	tokenParts := strings.Split(authorization, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token harus dalam format 'Bearer <token>'")
	}

	token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token salah")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Kesalahan dalam token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Kesalahan user ID dalam token claims")
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID salah / tidak ditemukan")
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "Kamu tidak berhak untuk menghapus akun ini")
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User tidak ditemukan")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses menghapus akun",
	})
}

func UpdateUserController(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		authorization = c.QueryParam("token")
	}

	if authorization == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token dibutuhkan")
	}

	tokenParts := strings.Split(authorization, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token harus dalam format 'Bearer <token>'")
	}

	token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Kesalahan authorization token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Kesalahan token claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Kesalahan user ID dalam token claims")
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if int(userID) != id {
		return echo.NewHTTPError(http.StatusForbidden, "Anda tidak di izinkan untuk mengubah data ini")
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User tidak ditemukan")
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := config.DB.Save(&existingUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses mengupdate data user",
		"user":    existingUser,
	})
}
