/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package createUser

// Import Librerary
// Librerary import
import (
	"backend/repository/query"
	"backend/repository/config"
	"fmt"
)

// Global Vars
var (
	_query_insert = config.ConfigMapSql()["create_user"]
	instanceQuery = query.GeneralQuery{}
)

/**  Create User Class */
type CreateUser struct {
	email_ string 
	password_ string 
	nickname_ string
	first_name_ string
	last_name_ string
	profile_picture_ string
}

/**
  * Insert New User
*/
func (p CreateUser) InsertNewUser(data ...*string) (map[string]interface{} , error) {
	// Execute the query
	response := map[string]interface{}{
		"status": false,
	}
	var is_error error
	
	result, is_error := query.Exec(instanceQuery, _query_insert,convertToStringInterfaceSlice(data)...)
	fmt.Println("Result :", result);

	// If error
	if is_error == nil {
		response["status"] = true
	}

	// Return result
	return response, is_error
}

func convertToStringInterfaceSlice(data []*string) []interface{} {
    result := make([]interface{}, len(data))
    for i, v := range data {
        result[i] = *v
    }
	
    return result
}