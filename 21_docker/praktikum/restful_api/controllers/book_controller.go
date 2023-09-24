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

func GetAllBooksController(c echo.Context) error {
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

	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
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

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
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

	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
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

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingBook models.Book
	if err := config.DB.First(&existingBook, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Publisher = book.Publisher

	if err := config.DB.Save(&existingBook).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    existingBook,
	})
}

func DeleteBookController(c echo.Context) error {
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

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}
