/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/

package followerControls;

// Librerary import
import (
	"github.com/gin-gonic/gin"
	"backend/repository/services/followers"
	"backend/src/services/httpControl"
	"backend/src/services/utils"
	"fmt"
)

// Vars
var (
	instanceRequest = httpControl.HttpRequestControl{}
	instanceFollowers = followers.FollowersGeneral{}
)

/**
  * Control of new follower.
*/
func InsertFollower(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	params_ := httpControl.GetBodyRequest(instanceRequest)
	requestParamns := utils.StateDefaultReq();

	if len(params_) == 2 {
		_response, err := followers.InsertFollower(instanceFollowers, utils.ConvertDataToArrayInt(params_)...);
		if err == nil {
			requestParamns["data"] = _response;
			requestParamns["status"] = 200;
			requestParamns["message_default"] = "Follower Inserted Success";
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
		fmt.Println(_response)
	}

	utils.ResponseControlGeneral(c, requestParamns)
}

/**
  * Control of new follower.
*/
func UpdateFollower(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	params_ := httpControl.GetBodyRequest(instanceRequest)
	requestParamns := utils.StateDefaultReq();

	if len(params_) == 3 {
		params_["expectedKeys"] = []string{"user_account", "follower_id" , "active"}
		_response, err := followers.UpdateFollower(instanceFollowers, utils.ConvertDataToArrayInterface(params_)...);
		if err == nil {
			requestParamns["data"] = _response;
			requestParamns["status"] = 200;
			requestParamns["message_default"] = "Update Success Follower";
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
		fmt.Println(_response)
	}

	utils.ResponseControlGeneral(c, requestParamns)
}