package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"17_ORM_and_Code_Structure/praktikum/soalEksplorasi_1-1/models"
)

// CreateBlogController membuat data blog baru
// CreateBlogController membuat data blog baru
func CreateBlogController(c echo.Context) error {
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validasi apakah UserID yang diberikan ada dalam tabel users
	var user models.User
	if err := DB.First(&user, blog.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid UserID")
	}

	if err := DB.Create(blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// GetAllBlogsController mendapatkan semua data blog
func GetAllBlogsController(c echo.Context) error {
	var blogs []models.Blog
	if err := DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// GetBlogController melihat rincian data blog berdasarkan ID
func GetBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var blog models.Blog
	if err := DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"blog":    blog,
	})
}

// UpdateBlogController mengubah data blog berdasarkan ID
func UpdateBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingBlog models.Blog
	if err := DB.First(&existingBlog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	existingBlog.Title = blog.Title
	existingBlog.Content = blog.Content

	if err := DB.Save(&existingBlog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    existingBlog,
	})
}

// DeleteBlogController menghapus data blog berdasarkan ID
func DeleteBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var blog models.Blog
	if err := DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	if err := DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}
