package controller

import (
	"gorm.io/gorm"
)

type UserController struct {
	db        *gorm.DB
	secretKey string
}

func NewUserController(db *gorm.DB, secretKey string) *UserController {
	return &UserController{
		db:        db,
		secretKey: secretKey,
	}
}
