/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package utils;

import (
	_ "regexp"
	"github.com/gin-gonic/gin"
)

/**
  * Conver data of request to []*string
  * Todo: Add Validation later
*/
func ConvertDataToArrayString(p map[string]interface{}) []*string {
	var response []*string
	expectedKeys := []string{"email", "password", "nickname", "firstname", "lastname", "img"}
	if len(p) == 6 && CheckExpectedKeys(p,expectedKeys) {
	   for _, key := range expectedKeys {
		 value := p[key].(string)
		 response = append(response, &value)
	   }
	}

	return response;
}

/**
  * Conver data of request to []*int
  * Todo: Add Validation later
*/
func ConvertDataToArrayInt(p map[string]interface{}) []*int {
	var response []*int
	expectedKeys := []string{"user_account", "follower_id"}

	if CheckExpectedKeys(p,expectedKeys) {
	   for _, key := range expectedKeys {
		 if value, ok := p[key].(float64); ok {
			valueInt := int(value)
		 	response = append(response, &valueInt)
		 } 
	   }
	}
	
	return response;
}

/**
  * Conver data of request to []*interface{} 
  * Todo: Add Validation later
*/
func ConvertDataToArrayInterface(p map[string]interface{}) []*interface{} {
	var response []*interface{}
	expectedKeys, ok := p["expectedKeys"].([]string)

	if ok && CheckExpectedKeys(p,expectedKeys) {
	   for _, key := range expectedKeys {
		 if value, ok := p[key]; ok {
		 	response = append(response, &value)
		 } 
	   }
	}
	
	return response;
}

/** 
  * Validate email
*/
// func validateEmail(email string) bool {
// 	re := regexp.MustCompile(`[\w\.-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9.-]+`);
// 	return re.MatchString(email)	
// }

// Check keys of map
func CheckExpectedKeys(m map[string]interface{}, expectedKeys []string) bool {
	haveAllKeys := true;
    for _, key := range expectedKeys {
        if _, ok := m[key]; !ok {
            haveAllKeys =  false
        }
    }

    return haveAllKeys;
}

/**
  * Return res to client
*/
func ResponseControlGeneral(c *gin.Context , args ...map[string]interface{}) {
	// Instance data of the request
	response := args[0]
	_status := response["status"].(int)
	_res := gin.H{
		"data": response["data"],
		"status": response["statusReq"], 
		"message": response["message_default"],
	}

	c.JSON(_status, _res);
}

/**
  * Data default to res 
*/
func StateDefaultReq() map[string]interface{} {
	return map[string]interface{}{
		"data": nil,
		"statusReq": false, 
		"message_default": "Error in request",	
		"status": 500,
	}
}