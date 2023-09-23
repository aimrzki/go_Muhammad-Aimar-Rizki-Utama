package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"ORM_and_Code_Structure_Full/config"
	"ORM_and_Code_Structure_Full/models"
)

func CreateBlogController(c echo.Context) error {
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := config.DB.First(&user, blog.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "UserId Salah")
	}

	if err := config.DB.Create(blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userName := user.Name
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil membuat blog",
		"blog": map[string]interface{}{
			"title":   blog.Title,
			"content": blog.Content,
			"user": map[string]interface{}{
				"name": userName,
			},
		},
	})
}

// GetAllBlogsController mendapatkan semua data blog
func GetAllBlogsController(c echo.Context) error {
	var blogs []models.Blog
	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for i, blog := range blogs {
		var user models.User
		if err := config.DB.First(&user, blog.UserID).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Gagal mengambil data user")
		}
		blogs[i].User = models.User{Name: user.Name}
	}

	var responseBlogs []map[string]interface{}
	for _, blog := range blogs {
		user := map[string]interface{}{
			"name": blog.User.Name,
		}

		responseBlogs = append(responseBlogs, map[string]interface{}{
			"ID":      blog.ID,
			"Title":   blog.Title,
			"Content": blog.Content,
			"User":    user,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses mengambil semua data blog",
		"blogs":   responseBlogs,
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
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses mendapatkan blog",
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
	if err := config.DB.First(&existingBlog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog tidak ditemukan")
	}

	existingBlog.Title = blog.Title
	existingBlog.Content = blog.Content

	if err := config.DB.Save(&existingBlog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil update blog",
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
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	if err := config.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses menghapus blog",
	})
}
