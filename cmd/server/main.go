package main

import (
	"fmt"
	transportHTTP "github.com/i-ms/klaar-backend/internal/transport/http"
	"net/http"
)

// App - struct containing information for things like
// pointer to database connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our app")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

// main -  Entry point for execution of Program
func main() {
	fmt.Println("Klaar REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting REST API")
		fmt.Println(err)
	}
}
