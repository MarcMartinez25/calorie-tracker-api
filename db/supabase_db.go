package db

import (
	"calorie-tracker-api/interfaces"
	"calorie-tracker-api/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

type SupabaseDB struct {
	client *supabase.Client
}

// NewSupabaseDB creates a new Supabase database connection
func NewSupabaseDB() (interfaces.Database, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	client, err := supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
	if err != nil {
		return nil, err
	}

	return &SupabaseDB{
		client: client,
	}, nil
}

// GetFoodEntryByID retrieves a single food entry by its ID
func (db *SupabaseDB) GetFoodEntryByID(id int) (*models.FoodEntry, error) {
	fmt.Printf("Fetching food entry with ID: %d\n", id)

	data, count, err := db.client.From("food_entries").
		Select("*", "exact", false).
		Eq("id", fmt.Sprintf("%d", id)).
		Execute()

	if err != nil {
		fmt.Printf("Error fetching food entry: %v\n", err)
		return nil, err
	}

	if count == 0 {
		return nil, nil // No entry found
	}

	var entries []models.FoodEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		fmt.Printf("Error unmarshaling data: %v\n", err)
		return nil, err
	}

	if len(entries) == 0 {
		return nil, nil // No entry found
	}

	return &entries[0], nil
}

// GetFoodEntriesByUserID retrieves all food entries for a specific user
func (db *SupabaseDB) GetFoodEntriesByUserID(userID int) ([]models.FoodEntry, error) {
	fmt.Printf("Fetching entries for user_id: %d\n", userID)

	data, count, err := db.client.From("food_entries").
		Select("*", "exact", false).
		Eq("user_id", fmt.Sprintf("%d", userID)).
		Execute()

	if err != nil {
		fmt.Printf("Error fetching entries: %v\n", err)
		return nil, err
	}

	fmt.Printf("Raw response data: %s\n", string(data))
	fmt.Printf("Count: %d\n", count)

	var entries []models.FoodEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		fmt.Printf("Error unmarshaling data: %v\n", err)
		return nil, err
	}

	fmt.Printf("Parsed entries: %+v\n", entries)
	return entries, nil
}
