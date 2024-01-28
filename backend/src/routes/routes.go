/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/
package routes

// Librerary import
import (
	"backend/src/controls"
    "backend/src/interfaces/httpRequest"
    "fmt"
)

/* Router type declaration */
func InitializeRoutes(router httpRequest.HttpRouter , allRoutes map[string]interface{}) {
   if allRoutes != nil || len(allRoutes) != 0 {
      for key := range allRoutes {
         router.GET(fmt.Sprintf("%s/user-email", key), httpRequest.HandleSync(controls.GetUserByEmail))
         router.POST(fmt.Sprintf("%s/user", key), httpRequest.HandleSync(controls.InsertNewUser))
      }
   }
}
