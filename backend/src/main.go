/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/
package main

// Librerary import
import (
	"github.com/gin-gonic/gin"
	"backend/repository/services/defaultRoutes"
	"backend/src/routes"
	"backend/db"
	"fmt"
	"sync"
)

// Constants
var (
	instanceRoutes = defaultRoutes.GeneralRoutes{}
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
		router := gin.Default()
		// Create a group of routes
		apiGroup := router.Group("/api")
		// Init all routes
		routes.InitializeRoutes(apiGroup, generateRoutes)
		router.Run(":4700")
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
	response, err := defaultRoutes.GetAllRoutes(instanceRoutes)
	if err != nil {
		fmt.Println("Error in get routes", err)
		response = make(map[string]interface{})
	}

	return response
}