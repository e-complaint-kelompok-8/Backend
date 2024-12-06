package config

import (
	"capstone/entities"
	"capstone/repositories/models"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&entities.Admin{},
		&models.Category{},
		&models.Complaint{},
		&models.ComplaintPhoto{},
		&models.News{},
		&models.Comment{},
		&models.Feedback{},
		// &entities.AIResponse{},
		// &entities.ChatMessage{},
		// &entities.ImportLog{},
		// &entities.Notification{},
	)
	log.Println("Migration completed successfully!")
}
