package main

import (
	"capstone/routes"
    "capstone/config"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
    // Jalankan migrasi database
    config.RunMigrations()
    
    // Membuat instance Echo
	e := echo.New()

	// Mendaftarkan routes
	routes.RegisterRoutes(e)

	// Menjalankan server pada port 8080
	log.Println("Server starting on port 8080...")
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}