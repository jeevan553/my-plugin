package main

import "fmt"

// Exec function now accepts strings instead of Args struct
func Exec(username, password, otherParam string) {
	fmt.Println("Hello from the plugin!")
	fmt.Println("Received Username:", username)
	fmt.Println("Received Password:", password)
	fmt.Println("Other Parameter:", otherParam)
}
