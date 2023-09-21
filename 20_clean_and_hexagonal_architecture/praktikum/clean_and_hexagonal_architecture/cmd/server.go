package cmd

import (
	"clean_and_hexagonal_architecture/internal/config"
	"clean_and_hexagonal_architecture/internal/controller"
	"log"

	"github.com/labstack/echo/v4"
)

func Start() error {
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Perform database migrations
	if err := config.MigrateDB(db); err != nil {
		log.Fatal(err)
	}

	// Create a new Echo instance
	e := echo.New()

	// Initialize dependencies
	userController := controller.NewUserController(db, "your-secret-key")

	// Define the routes
	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
	e.POST("/login", userController.LoginUser)

	// Start the server
	return e.Start(":8080")
}
