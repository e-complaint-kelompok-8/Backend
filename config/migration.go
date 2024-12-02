package config

import (
	"capstone/entities"
	"capstone/repositories/models"
	"log"

	"gorm.io/gorm"
)

// RunMigrations is used to perform database migrations
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&entities.Admin{},
		&models.User{},
		// &entities.AIResponse{},
		&models.Category{},
		// &entities.ChatMessage{},
		// &entities.Commentar{},
		&models.Complaint{},
		// &entities.Feedback{},
		// &entities.ImportLog{},
		// &entities.News{},
		// &entities.Notification{},
		&models.ComplaintPhoto{},
	)
	log.Println("Migration completed successfully!")
}
