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
)

var (
	instanceRequest = httpControl.HttpRequestControl{}
	instDataUser = dataUserEmail.DataUserEmail{}
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
			fmt.Println(err)
		}
	}

	c.JSON(_status, _res)
}