/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/

package userControls;

// Librerary import
import (
	"github.com/gin-gonic/gin"
	"backend/src/services/httpControl"
	"backend/repository/services/createUser"
	"backend/repository/services/userUtils"
	"backend/src/services/utils"
	"fmt"
)

// Vars
var (
	instanceRequest = httpControl.HttpRequestControl{}
	instanceCreateuser = createUser.CreateUser{}
	instanceUser = userUtils.UserGeneralUtils{}
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
		_response, err := userUtils.GetDataUserEmail(instanceUser,email_.(string));
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
 * Return Follower per user
*/
func GetFollowerUser(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	user_account, ok := (httpControl.GetBodyRequest(instanceRequest))["user_account"]
	requestParamns := utils.StateDefaultReq();

	if ok {
		user_account64 := user_account.(float64)
		fmt.Println(int(user_account64))
		_response, err := userUtils.GetDataFollwer(instanceUser,int(user_account64));
		if err == nil {
			requestParamns["status"] = 200;
			if len(_response) > 0 {
				requestParamns["data"] = _response;
				requestParamns["message_default"] = "Followers Found";
			} else {
				requestParamns["message_default"] = "No have Followers";
			}
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
	}

	utils.ResponseControlGeneral(c, requestParamns)
}

/**
 * Return Follower per user
*/
func GetSubscriptionUser(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	user_account, ok 	:= (httpControl.GetBodyRequest(instanceRequest))["user_account"]
	requestParamns := utils.StateDefaultReq();

	if ok {
		user_account64 := user_account.(float64)
		_response, err := userUtils.GetDataSubscription(instanceUser,int(user_account64));
		if err == nil {
			requestParamns["status"] = 200;
			if len(_response) > 0 {
				requestParamns["data"] = _response;
				requestParamns["message_default"] = "Subscription Found";
			} else {
				requestParamns["message_default"] = "No have Subscription";
			}
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
	}

	utils.ResponseControlGeneral(c, requestParamns)
}

/**
 * Return Follower per user
*/
func GetDashboardUser(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	user_account, ok 	:= (httpControl.GetBodyRequest(instanceRequest))["user_account"]
	requestParamns := utils.StateDefaultReq();

	if ok {
		user_account64 := user_account.(float64)
		_response, err := userUtils.GetDataDashboard(instanceUser,int(user_account64));
		if err == nil {
			requestParamns["status"] = 200;
			if len(_response) > 0 {
				requestParamns["data"] = _response;
				requestParamns["message_default"] = "Dashboard Found";
			} else {
				requestParamns["message_default"] = "No have Dashboard";
			}
		} else {
			fmt.Println(err);
			requestParamns["message_default"] = err;
		}
	}

	utils.ResponseControlGeneral(c, requestParamns)
}