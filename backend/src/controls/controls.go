/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/

package controls;

// Librerary import
import (
	"github.com/gin-gonic/gin"
	"backend/src/services/httpControl"
	"backend/repository/services/dataUserEmail"
	"backend/repository/services/createUser"
	"backend/repository/services/followers"
	"backend/src/services/utils"
	"fmt"
)

// Vars
var (
	instanceRequest = httpControl.HttpRequestControl{}
	instDataUser = dataUserEmail.DataUserEmail{}
	instanceCreateuser = createUser.CreateUser{}
	instanceFollwers = followers.FollersGeneral{}
)

/**
 * Return Data per email
*/
func GetUserByEmail(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	email_, ok 	:= (httpControl.GetBodyRequest(instanceRequest))["email"]
	requestParamns := utils.StateDefaultReq();

	if ok {
		_response, err := dataUserEmail.GetDataUserEmail(instDataUser,email_.(string));
		if err == nil {
			requestParamns["status"] = 200;
			if len(_response) > 0 {
				requestParamns["data"] = _response;
				requestParamns["message_default"] = "user found";
			} else {
				requestParamns["message_default"] = "user not found";
			}
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
	}

	utils.ResponseControlGeneral(c, requestParamns)
}

/**
  * Control of new user.
*/
func InsertNewUser(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	params_ := (httpControl.GetBodyRequest(instanceRequest))
	requestParamns := utils.StateDefaultReq();
	
	if len(params_) == 6 {
		_response, err := createUser.InsertNewUser(instanceCreateuser, utils.ConvertDataToArrayString(params_)...);
		if err == nil {
			requestParamns["data"] = _response;
			requestParamns["status"] = 200;
			requestParamns["message_default"] = "User Created";
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
func InsertFollower(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	params_ := httpControl.GetBodyRequest(instanceRequest)
	requestParamns := utils.StateDefaultReq();

	if len(params_) == 2 {
		_response, err := followers.InsertFollower(instanceFollwers, utils.ConvertDataToArrayInt(params_)...);
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