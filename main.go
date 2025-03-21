package main

import (
	"log"
	"net/http"
	"os"

	"calorie-tracker-api/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize database
	if err := db.Initialize(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize router
	r := gin.Default()

	// Routes
	r.GET("/api/users/:userId/food-entries", getFoodEntries)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// getFoodEntries handles GET requests for food entries by user ID
func getFoodEntries(c *gin.Context) {
	userId := c.Param("userId")

	entries, err := db.GetFoodEntriesByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch food entries",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"entries": entries,
	})
}
