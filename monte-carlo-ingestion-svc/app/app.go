package app

type App interface {
	// Load is responsible for loading the necessary configurations and dependencies of the application.
	Load()

	// Start is responsible for starting the application and its associated components.
	Start()
}
