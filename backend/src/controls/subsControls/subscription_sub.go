/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/

package subsControls;

// Librerary import
import (
	"github.com/gin-gonic/gin"
	"backend/src/services/httpControl"
	"backend/repository/services/subscription"
	"backend/src/services/utils"
	"fmt"
)

// Vars
var (
	instanceRequest = httpControl.HttpRequestControl{}
	instanceSubscription = subscription.Subscription{}
)

/**
  * Control of update Subscription.
*/
func UpdateSubscription(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	params_ := httpControl.GetBodyRequest(instanceRequest)
	requestParamns := utils.StateDefaultReq();

	if len(params_) == 6 {
		params_["expectedKeys"] = []string{
			"user_account", 
			"follower_id" , 
			"json",
			"active_sub",
			"start_sub_p",
			"finish_sub",
		}
		_response, err := subscription.UpdateSubscription(instanceSubscription, utils.ConvertDataToArrayInterface(params_)...);
		if err == nil {
			requestParamns["data"] = _response;
			requestParamns["status"] = 200;
			requestParamns["message_default"] = "Update Success Subscription";
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
		fmt.Println(_response)
	}

	utils.ResponseControlGeneral(c, requestParamns)
}

/**
  * Control of new Subscription.
*/
func InsertSubscription(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	params_ := httpControl.GetBodyRequest(instanceRequest)
	requestParamns := utils.StateDefaultReq();

	if len(params_) == 6 {
		params_["expectedKeys"] = []string{
			"user_account", 
			"follower_id" , 
			"json",
			"active_sub",
			"start_sub_p",
			"finish_sub",
		}
		_response, err := subscription.InsertSubscription(instanceSubscription, utils.ConvertDataToArrayInterface(params_)...);
		if err == nil {
			requestParamns["data"] = _response;
			requestParamns["status"] = 200;
			requestParamns["message_default"] = "Insert Success Subscription";
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
		fmt.Println(_response)
	}

	utils.ResponseControlGeneral(c, requestParamns)
}