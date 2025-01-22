// Package main is the entry point for our Go program
// Every executable Go program must have a main package
package main

// Import section - these are the packages we'll use
import (
	"fmt" // fmt is used for formatting and printing
	// log package for logging errors
	"time"                // time package for working with dates and time
	"your-module-name/db" // Import the db package
)

// main is the entry point function for our program
// It must be named "main" and take no parameters
func main() {
	// Basic variable declaration and initialization
	// In Go, you can use := for short variable declaration
	name := "Learner"

	// Traditional variable declaration
	var age int = 25

	// Printing to console
	// Printf allows formatted printing
	fmt.Printf("Hello, %s! You are %d years old\n", name, age)

	// Demonstrating a slice (dynamic array)
	// [] indicates a slice, string is the type
	fruits := []string{"apple", "banana", "orange"}

	// For loop example
	// range is used to iterate over slices, arrays, maps, etc.
	fmt.Println("\nLet's learn about fruits:")
	for index, fruit := range fruits {
		fmt.Printf("Fruit %d: %s\n", index+1, fruit)
	}

	// Demonstrating if-else condition
	currentHour := time.Now().Hour()
	if currentHour < 12 {
		fmt.Println("\nGood morning!")
	} else if currentHour < 18 {
		fmt.Println("\nGood afternoon!")
	} else {
		fmt.Println("\nGood evening!")
	}

	// Calling a custom function
	result := addNumbers(5, 3)
	fmt.Printf("\n5 + 3 = %d\n", result)

	// Demonstrating error handling
	result, err := divideNumbers(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 ÷ 2 = %d\n", result)
	}

	// Fetching data from the database
	db.FetchDataFromDB()
}

// Function declaration
// func functionName(parameter parameterType) returnType
func addNumbers(a int, b int) int {
	return a + b
}

// Function that returns multiple values (result and error)
// This is a common pattern in Go for error handling
func divideNumbers(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
