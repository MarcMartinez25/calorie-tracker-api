package db

import (
	"calorie-tracker-api/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/supabase-community/supabase-go"
)

var client *supabase.Client

// Initialize sets up the Supabase client
func Initialize() error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	var err error
	client, err = supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
	if err != nil {
		return err
	}

	return nil
}

// GetFoodEntriesByUserID retrieves all food entries for a specific user
func GetFoodEntriesByUserID(userID string) ([]models.FoodEntry, error) {
	var entries []models.FoodEntry

	// Convert string userID to int for validation
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	fmt.Printf("Fetching entries for user_id: %d\n", userIDInt)

	// Temporarily removing the filter to see if we can get any data
	data, count, err := client.From("food_entries").
		Select("*", "exact", false).
		Eq("user_id", userID).
		Execute()

	if err != nil {
		fmt.Printf("Error fetching entries: %v\n", err)
		return nil, err
	}

	fmt.Printf("Raw response data: %s\n", string(data))
	fmt.Printf("Count: %d\n", count)

	if err := json.Unmarshal(data, &entries); err != nil {
		fmt.Printf("Error unmarshaling data: %v\n", err)
		return nil, err
	}

	fmt.Printf("Parsed entries: %+v\n", entries)
	return entries, nil
}
