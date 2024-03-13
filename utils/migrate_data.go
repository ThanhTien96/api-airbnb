package utils

import (
	"log"

	"github.com/ThanhTien96/airbnb/models"
	"gorm.io/gorm"
)

func MigrateData(db *gorm.DB) error {
	log.Println("Migrating data...")
	err := db.AutoMigrate(
		&models.UserRole{},
		&models.User{},
	)
	if err != nil {
		return err
	}
	return nil
}