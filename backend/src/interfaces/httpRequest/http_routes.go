/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/
package httpRequest

// Librerary import
import (
  "github.com/gin-gonic/gin"
  "sync"
)

// Mutex
var mutex = &sync.Mutex{}

/* Interface declaration */

/**
  *
  * @decorator: HttpRouter
  * @description: This interface is used to define the methods of the router
  * @method: GET
  * @method: POST
  * @method: PUT
  * @method: DELETE
*/
type HttpRouter interface {
	// Method GET
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	// Method POST
	POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	// Method PUT
	PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

	// Method DELETE
	DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}

/**
  * Sync Handlers functions
  *
  * @param: gin.HandlerFunc
**/
func HandleSync(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
	  h(c)
	}
}