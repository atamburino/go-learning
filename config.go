// Package main is the entry point for our Go program
package main

import (
	"os" // os package for environment variables
)

// Function to get the database connection string from environment variables
func getDBConnectionString() string {
	return os.Getenv("DATABASE_URL")
}
