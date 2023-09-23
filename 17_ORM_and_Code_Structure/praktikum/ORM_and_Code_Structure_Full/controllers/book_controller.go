package controllers

import (
	"net/http"
	"strconv"

	"ORM_and_Code_Structure_Full/config"
	"ORM_and_Code_Structure_Full/models"
	"github.com/labstack/echo"
)

// GetAllBooksController mengembalikan semua data buku
func GetAllBooksController(c echo.Context) error {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses mendapatkan semua data buku",
		"books":   books,
	})
}

// GetBookController mengembalikan data buku berdasarkan ID
func GetBookController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID salah")
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Buku tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses mendapatkan buku dengan ID",
		"book":    book,
	})
}

// CreateBookController membuat buku baru
func CreateBookController(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses membuat buku baru",
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
	if err := config.DB.First(&existingBook, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Buku tidak ditemukan")
	}

	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Publisher = book.Publisher

	if err := config.DB.Save(&existingBook).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses mengupdate data buku",
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
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Buku tidak ditemukan")
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil menghapus buku",
	})
}
