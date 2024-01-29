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
func InitializeRoutes(router httpRequest.HttpRouter , allRoutes []map[string]interface{}) {
   if allRoutes != nil || len(allRoutes) != 0 {
      for _, route := range allRoutes {
         router.GET(fmt.Sprintf("%s/user-email", route["route_"]), httpRequest.HandleSync(controls.GetUserByEmail))
         router.POST(fmt.Sprintf("%s/user", route["route_"]), httpRequest.HandleSync(controls.InsertNewUser))
         router.POST(fmt.Sprintf("%s/follower", route["route_"]), httpRequest.HandleSync(controls.InsertFollower))
      }
   }
}
