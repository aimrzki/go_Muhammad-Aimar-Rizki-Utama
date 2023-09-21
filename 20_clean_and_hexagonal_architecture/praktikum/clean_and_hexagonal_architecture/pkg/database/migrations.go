package database

import (
	"clean_and_hexagonal_architecture/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
	)
}
