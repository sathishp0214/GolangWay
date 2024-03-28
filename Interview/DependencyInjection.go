package main

import (
	"fmt"
	"net/http"
)

/*
Dependency injection - Passing dependencies as argument. Passing interface implemented different struct types as argument into function.

Example:
func NewApplication(logger Logger) *Application { //Here we passing different "Logger" interface implemented structs dynamically.
	return &Application{logger: logger}
}
*/

// Logger interface represents a generic logger with a single method Log.
type Logger interface {
	Log(message string)
}

// FileLogger is an implementation of the Logger interface that logs messages to a file.
type FileLogger struct {
	// Additional fields and configurations for the file logger can be added here.
}

// Log logs the given message to a file.
func (fl *FileLogger) Log(message string) {
	// Code to write the message to a file.
	fmt.Println("[File Logger] " + message)
}

// ConsoleLogger is an implementation of the Logger interface that logs messages to the console.
type ConsoleLogger struct {
	// Additional fields and configurations for the console logger can be added here.
}

// Log logs the given message to the console.
func (cl *ConsoleLogger) Log(message string) {
	// Code to print the message to the console.
	fmt.Println("[Console Logger] " + message)
}

// Application represents our main application that depends on a Logger.
type Application struct {
	logger Logger // Dependency injection of the logger.
}

// NewApplication creates a new instance of Application with the given logger dependency.
func NewApplication(logger Logger) *Application { //Here we passing different "Logger" interface implemented structs dynamically.
	return &Application{logger: logger}
}

// Run starts the application and logs a message.
func (app *Application) Run() {
	// Application logic goes here.
	app.logger.Log("Application is running.")
}

func main() {
	// Create a FileLogger instance for dependency injection.
	fileLogger := &FileLogger{}
	// Create a ConsoleLogger instance for dependency injection.
	consoleLogger := &ConsoleLogger{}
	// Create a new Application instance with FileLogger dependency.
	appWithFileLogger := NewApplication(fileLogger)
	appWithFileLogger.Run()
	// Create a new Application instance with ConsoleLogger dependency.
	appWithConsoleLogger := NewApplication(consoleLogger)
	appWithConsoleLogger.Run()
}

// ------------------------------------------------------------------------------
// Router defines the interface for a router.
type Router interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

// MyRouter is an example implementation of the Router interface.
type MyRouter struct{}

func (r *MyRouter) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}

func (r *MyRouter) ServeHTTP(http.ResponseWriter, *http.Request) {
}

// LoggerMiddleware is an example middleware.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a router instance
	router := &MyRouter{} //MyRouter is our mock struct, Implments our "Router" mock interface,

	// Attach middleware
	http.Handle("/", LoggerMiddleware(router))

	// These below function calls our mock struct's HandleFunc(), There we can do our unit testing.
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to our website!")
	})

	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About Us")
	})

	router.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contact Us")
	})

	// Start the server
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
