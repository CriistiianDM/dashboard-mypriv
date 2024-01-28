/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package utils;

import (
	"regexp"
)

/**
  * Conver data of request to []*string
  * Todo: Add Validation later
*/
func ConvertDataToArrayCreateUser(p map[string]interface{}) []*string {
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