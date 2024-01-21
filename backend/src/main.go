/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package main

// Librerary import
import (
	_ "github.com/gin-gonic/gin"
	"backend/db"
	"backend/repository/generalSettings"
	"fmt"
	"sync"
)

// Constants
var (
	instanceRoutes = generalSettings.GeneralSettings{}
)

/**
  * Main function
*/
func main() {
	_executeRun();
	_initServer();
}

/**
  * Initialize the server
*/
func _initServer() {
	generateRoutes := _initRoutes()

	if len(generateRoutes) != 0 {
		//Initialize the route server
	}
}

/**
  * Execute all goroutines
*/
func _executeRun() {
	var wg sync.WaitGroup
		wg.Add(1)
		go db.Connect(&wg);
	wg.Wait()
}

/**
  * Initialize the routes
*/
func _initRoutes() map[string]interface{} {
	response, err := generalSettings.GetAllRoutes(instanceRoutes)
	if err != nil {
		fmt.Println("Error in get routes", err)
		response = make(map[string]interface{})
	}

	return response
}