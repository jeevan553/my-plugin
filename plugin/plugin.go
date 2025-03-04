// plugin/plugin.go
package main

import (
	"fmt"
	"os"
)

// Exec is the function that the main program will use from the plugin
func Exec() {
	// Get environment variables (for demonstration)
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	otherParam := os.Getenv("OTHER_PARAM")

	// Print values to show plugin works
	fmt.Println("Hello from the plugin!")
	fmt.Println("Received Username:", username)
	fmt.Println("Received Password:", password)
	fmt.Println("Other Parameter", otherParam)
}
