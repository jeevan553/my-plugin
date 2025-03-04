// main.go
package main

import (
	"log"
	"plugin"
)

func main() {
	// Set environment variables
	//os.Setenv("USERNAME", "user123")
	//os.Setenv("PASSWORD", "password123")
	//os.Setenv("OTHER_PARAM", "some_value")

	// Open the plugin
	plug, err := plugin.Open("plugin.so")
	if err != nil {
		log.Fatal(err)
	}

	// Lookup the Exec function in the plugin
	execFunc, err := plug.Lookup("Exec")
	if err != nil {
		log.Fatal(err)
	}

	// Assert the function type
	hello, ok := execFunc.(func())
	if !ok {
		log.Fatal("Plugin has no 'Exec' function with the correct signature")
	}

	// Call the Exec function from the plugin
	hello()

	// Output: "Hello from the plugin!"
}
