// Package main is the entry point for our Go program
package main

import (
	"database/sql" // Importing the database/sql package
	"fmt"          // fmt package for printing
	"log"          // log package for logging errors

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Function to fetch data from the database
func fetchDataFromDB() {
	// Database connection string
	connStr := getDBConnectionString() // Fetch connection string from config
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Querying the database
	rows, err := db.Query("SELECT id, name FROM users") // Update with your query
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterating through the results
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
