package main

import (
	"log"
	"os"

	"calorie-tracker-api/controllers"
	"calorie-tracker-api/db"
	"calorie-tracker-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize database
	database, err := db.Initialize()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize services
	foodEntryService := services.NewFoodEntryService(database)

	// Initialize controllers
	foodEntriesController := controllers.NewFoodEntriesController(foodEntryService)

	// Initialize router
	router := gin.Default()

	// Routes
	router.GET("/api/users/:userId/food-entries", foodEntriesController.GetFoodEntries)
	router.GET("/api/food-entries/:id", foodEntriesController.GetFoodEntry)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
