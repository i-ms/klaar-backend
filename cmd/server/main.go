package main

import "fmt"

// App - struct containing information for things like
// pointer to database connections
type App struct{

}

// Run - sets up our application
func (app *App)Run()error{
	fmt.Println("Setting up our app")
	return nil
}

// main -  Entry point for execution of Program
func main(){
	fmt.Println("Klaar REST API")
	app:=App{}
	if err:=app.Run();err!=nil{
		fmt.Println("Error starting REST API")
		fmt.Println(err)
	}
}
