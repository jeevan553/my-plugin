package main

import (
	"log"
	"os"
	"plugin"
	"reflect"
)

func main() {
	// Open the plugin
	plug, err := plugin.Open("plugin.so")
	if err != nil {
		log.Fatal(err)
	}

	// Lookup the Exec function
	execFunc, err := plug.Lookup("Exec")
	if err != nil {
		log.Fatal(err)
	}

	// Use reflection to call Exec dynamically
	exec := reflect.ValueOf(execFunc)

	// Ensure Exec is a function
	if exec.Kind() != reflect.Func {
		log.Fatal("Exec is not a function")
	}

	// Get environment variables
	args := []reflect.Value{
		reflect.ValueOf(os.Getenv("USERNAME")),
		reflect.ValueOf(os.Getenv("PASSWORD")),
		reflect.ValueOf(os.Getenv("OTHERPARAM")),
	}

	// Call the function dynamically
	exec.Call(args)
}
