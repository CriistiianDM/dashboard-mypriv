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
)

var (
	instanceRequest = httpControl.HttpRequestControl{}
)

/**
 * Return Data per email
*/
func GetUserByEmail(c *gin.Context) {
	// Instance data of the request
	instanceRequest.BodyRequest = c
	paramsBody := httpControl.GetBodyRequest(instanceRequest)
	email_, ok := paramsBody["email"]

	if ok {
		fmt.Println("text" , email_)
	}
	c.JSON(200, gin.H{"response": true , "message": "Product inserted"})
}