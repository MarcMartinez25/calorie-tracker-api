package interfaces

import "calorie-tracker-api/models"

// Database defines the interface for database operations
type Database interface {
	GetFoodEntryByID(id int) (*models.FoodEntry, error)
	GetFoodEntriesByUserID(userID int) ([]models.FoodEntry, error)
}
