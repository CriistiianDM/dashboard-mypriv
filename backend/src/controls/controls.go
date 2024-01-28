/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/

package controls;

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"backend/src/services/httpControl"
	"backend/repository/services/dataUserEmail"
	"backend/repository/services/createUser"
	"backend/src/services/utils"
	_ "encoding/json"
)

var (
	instanceRequest = httpControl.HttpRequestControl{}
	instDataUser = dataUserEmail.DataUserEmail{}
	instanceCreateuser = createUser.CreateUser{}
)

/**
 * Return Data per email
*/
func GetUserByEmail(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	email_, ok 	:= (httpControl.GetBodyRequest(instanceRequest))["email"]
    _status := 500
	_res := gin.H{"status": false , "message": "No found user or email no register"}

	if ok {
		_response, err := dataUserEmail.GetDataUserEmail(instDataUser,email_.(string));
		if err == nil {
			_res =  gin.H{"data": _response, "status": 200, "message": "user found"}
			_status = 200
		} else {
			fmt.Println(err);
			_res = gin.H{"status": false , "message": err}
		}
	}

	c.JSON(_status, _res);
}

/**
  * Control of new user.
*/
func InsertNewUser(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	params_ := (httpControl.GetBodyRequest(instanceRequest))
    _status := 500
	_res := gin.H{"status": false , "message": "No data complete for insert create user."}
	
	if len(params_) == 6 {
		_response, err := createUser.InsertNewUser(instanceCreateuser, utils.ConvertDataToArrayCreateUser(params_)...);
		if err == nil {
			_res =  gin.H{"data": _response, "status": 200, "message": "user found"}
			_status = 200
		} else {
			fmt.Println(err);
			_res = gin.H{"status": false , "message": err}
		}
		fmt.Println(_response)
	}

	c.JSON(_status, _res);
}