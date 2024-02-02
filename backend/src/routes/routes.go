/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/
package routes

// Librerary import
import (
   "backend/src/interfaces/httpRequest"
   "backend/src/controls/userControls"
   "backend/src/controls/followerControls"
   "backend/src/controls/subsControls"
   "fmt"
)

/* Router type declaration */
func InitializeRoutes(router httpRequest.HttpRouter , allRoutes []map[string]interface{}) {
   if allRoutes != nil || len(allRoutes) != 0 {
      for _, route := range allRoutes {
         // User
         router.POST(fmt.Sprintf("%s/user-email", route["route_"]), httpRequest.HandleSync(userControls.GetUserByEmail))
         router.GET(fmt.Sprintf("%s/user-follwer", route["route_"]), httpRequest.HandleSync(userControls.GetFollowerUser))
         router.GET(fmt.Sprintf("%s/user-subscription", route["route_"]), httpRequest.HandleSync(userControls.GetSubscriptionUser))
         router.GET(fmt.Sprintf("%s/user-dashboard", route["route_"]), httpRequest.HandleSync(userControls.GetDashboardUser))     
         router.POST(fmt.Sprintf("%s/user", route["route_"]), httpRequest.HandleSync(userControls.InsertNewUser))
         router.POST(fmt.Sprintf("%s/all-user", route["route_"]),httpRequest.HandleSync(userControls.GetDataAllUser)) 
         // Follower
         router.POST(fmt.Sprintf("%s/follower", route["route_"]), httpRequest.HandleSync(followerControls.InsertFollower))
         router.PUT(fmt.Sprintf("%s/follower", route["route_"]), httpRequest.HandleSync(followerControls.UpdateFollower))
         // Subscription
         router.POST(fmt.Sprintf("%s/subscription", route["route_"]), httpRequest.HandleSync(subsControls.InsertSubscription))
         router.PUT(fmt.Sprintf("%s/subscription", route["route_"]), httpRequest.HandleSync(subsControls.UpdateSubscription))
      }
   }
}
