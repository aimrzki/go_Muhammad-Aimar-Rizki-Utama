package config

import (
	"ORM_and_Code_Structure_Full/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	config := LoadConfig()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})
}
