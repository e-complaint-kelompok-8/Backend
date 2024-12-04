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
		&models.Category{},
		&models.Complaint{},
		&models.ComplaintPhoto{},
		&entities.Admin{},
		&models.News{},
		&models.Comment{},
		// &entities.AIResponse{},
		// &models.Category{},
		// &entities.ChatMessage{},
		// &entities.Commentar{},
		// &models.Complaint{},
		// &entities.Feedback{},
		// &entities.ImportLog{},
		// &entities.News{},
		// &entities.Notification{},
	)
	log.Println("Migration completed successfully!")
}
