package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	// Mendapatkan ID dari URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	// Cari user berdasarkan ID
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by ID",
				"user":     user,
			})
		}
	}

	// Jika user tidak ditemukan
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": "User not found",
	})
}

func DeleteUserController(c echo.Context) error {
	// Mendapatkan ID dari URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	// Cari user berdasarkan ID
	for i, user := range users {
		if user.Id == id {
			// Hapus user dari slice
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user",
			})
		}
	}

	// Jika user tidak ditemukan
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": "User not found",
	})
}

func UpdateUserController(c echo.Context) error {
	// Mendapatkan ID dari URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	// Cari user berdasarkan ID
	for i, user := range users {
		if user.Id == id {
			// Binding data update
			c.Bind(&user)
			users[i] = user
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user",
				"user":     user,
			})
		}
	}

	// Jika user tidak ditemukan
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": "User not found",
	})
}

func CreateUserController(c echo.Context) error {
	// Binding data user baru
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	// Routing untuk mengakses endpoint-endpoint CRUD
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users/add", CreateUserController)
	e.PUT("/users/update/:id", UpdateUserController)
	e.DELETE("/users/delete/:id", DeleteUserController)

	// Menjalankan server pada port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
