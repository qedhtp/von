package main

import (
	"fmt"
	"os"
)

// Function that accepts one string argument
func get_query(arg *string) {
	fmt.Println("Query argument:", arg)
}

func main() {
	// Check if a command-line argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command-line argument.")
		return
	}

	// Get the first argument from the command line (os.Args[1])
	arg := os.Args[1]

	// Pass the argument to the function
	get_query(&arg)
}