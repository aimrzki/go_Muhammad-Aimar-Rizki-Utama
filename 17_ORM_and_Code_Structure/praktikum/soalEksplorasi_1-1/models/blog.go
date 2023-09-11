package models

import (
	"17_ORM_and_Code_Structure/praktikum/soalEksplorasi_1-1/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
}

func InitDB() {
	config := config.LoadConfig()
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

type Blog struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Hubungan dengan model User
func (b *Blog) User() *User {
	var user User
	DB.Model(&b).Related(&user, "UserID")
	return &user
}
