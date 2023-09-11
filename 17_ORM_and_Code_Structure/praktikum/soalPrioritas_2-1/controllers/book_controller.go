package controllers

import (
	"net/http"
	"strconv"

	"17_ORM_and_Code_Structure/praktikum/soalPrioritas_2-1/models"
	"github.com/labstack/echo"
)

// GetAllBooksController mengembalikan semua data buku
func GetAllBooksController(c echo.Context) error {
	var books []models.Book
	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// GetBookController mengembalikan data buku berdasarkan ID
func GetBookController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var book models.Book
	if err := DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// CreateBookController membuat buku baru
func CreateBookController(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := DB.Create(book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// UpdateBookController mengubah data buku berdasarkan ID
func UpdateBookController(c echo.Context) error {
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
	if err := DB.First(&existingBook, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Publisher = book.Publisher

	if err := DB.Save(&existingBook).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    existingBook,
	})
}

// DeleteBookController menghapus data buku berdasarkan ID
func DeleteBookController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var book models.Book
	if err := DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	if err := DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}
