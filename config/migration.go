package config

import (
	"capstone/repositories/models"
	"capstone/entities"
	"log"

	"gorm.io/gorm"
)

// RunMigrations is used to perform database migrations
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&entities.Admin{},
		// &entities.AIResponse{},
		// &models.Category{},
		// &entities.ChatMessage{},
		// &entities.Commentar{},
		// &models.Complaint{},
		// &entities.Feedback{},
		// &entities.ImportLog{},
		// &entities.News{},
		// &entities.Notification{},
		// &models.ComplaintPhoto{},
		&models.User{},
	)
	log.Println("Migration completed successfully!")
}
