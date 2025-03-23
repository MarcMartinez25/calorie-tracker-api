package db

import "calorie-tracker-api/interfaces"

// Initialize sets up the database connection and returns a Database interface
func Initialize() (interfaces.Database, error) {
	return NewSupabaseDB()
}
