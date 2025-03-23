package services

import (
	"calorie-tracker-api/interfaces"
	"calorie-tracker-api/models"
	"fmt"
	"strconv"
)

type FoodEntryService struct {
	db interfaces.Database
}

func NewFoodEntryService(db interfaces.Database) *FoodEntryService {
	return &FoodEntryService{
		db: db,
	}
}

// GetByID retrieves a single food entry and handles business logic
func (s *FoodEntryService) GetByID(id string) (*models.FoodEntry, error) {
	// Validate ID
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// Additional business logic could go here
	// For example:
	// - Checking user permissions
	// - Data transformation
	// - Validation rules

	return s.db.GetFoodEntryByID(idInt)
}

// GetByUserID retrieves all food entries for a user and handles business logic
func (s *FoodEntryService) GetByUserID(userID string) ([]models.FoodEntry, error) {
	// Validate user ID
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %v", err)
	}

	// Additional business logic could go here
	// For example:
	// - Filtering entries by date
	// - Sorting entries
	// - Calculating daily totals
	// - Applying user preferences

	return s.db.GetFoodEntriesByUserID(userIDInt)
}
