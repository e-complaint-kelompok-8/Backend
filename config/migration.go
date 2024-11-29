package config

import (
	"capstone/entities"
	"log"
)

// RunMigrations is used to perform database migrations
func RunMigrations() {
	// Connect to the database
	db, err := ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform the migration
	log.Println("Running migrations...")
	err = db.AutoMigrate(
		&entities.Admin{},
		&entities.AIResponse{},
		&entities.Category{},
		&entities.ChatMessage{},
		&entities.Commentar{},
		&entities.Complaint{},
		&entities.Feedback{},
		&entities.ImportLog{},
		&entities.News{},
		&entities.Notification{},
		&entities.UrlPhoto{},
		&entities.User{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully!")
}