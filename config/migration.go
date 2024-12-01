package config

import (
	"capstone/repositories/models"
	"log"

	"gorm.io/gorm"
)

// RunMigrations is used to perform database migrations
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Complaint{},
		&models.ComplaintPhoto{},
		// &entities.Admin{},
		// &entities.AIResponse{},
		// &entities.ChatMessage{},
		// &entities.Commentar{},
		// &entities.Feedback{},
		// &entities.ImportLog{},
		// &entities.News{},
		// &entities.Notification{},
	)
	log.Println("Migration completed successfully!")
}
