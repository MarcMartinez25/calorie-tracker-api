package models

import "time"

// FoodEntry represents a single food entry in the database
type FoodEntry struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Calories  int       `json:"calories"`
	Protein   int       `json:"protein"`
	Carbs     int       `json:"carbs"`
	Fat       int       `json:"fat"`
	CreatedAt time.Time `json:"created_at"`
}
