package config

import (
	"capstone/repositories/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	// AutoMigrate tabel
	db.AutoMigrate(
		&models.Complaint{},
		&models.ComplaintPhoto{},
		&models.User{},
		&models.Admin{},
		&models.Category{},
		&models.News{},
		&models.Comment{},
		&models.Feedback{},
	)
	log.Println("Migration completed successfully!")
}
