package main

import (
	"monte-carlo-ingestion/wire"
)

// main is the entry point of the application.
func main() {
	// Initialize the application and start the server.
	server, err := wire.InitializeApplication()
	if err != nil {
		panic(err)
	}

	server.Load()
	server.Start()
}
