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
	"backend/repository/config"
	"backend/repository/services/utils"
)

// Global Vars
var (
	_query_insert = config.ConfigMapSql()["create_user"]
)

/**  Create User Class */
type CreateUser struct {}

/**
  * Insert New User
*/
func (p CreateUser) InsertNewUser(data ...*string) (map[string]interface{} , error) {
	// Return result
	return utils.GetDataExec(_query_insert,p.convertInterfaceSlice(data)...)
}

/**
  * Convert []*string to []interface{}
*/
func (p CreateUser) convertInterfaceSlice(data []*string) []interface{} {
    result := make([]interface{}, len(data))
    for i, v := range data {
        result[i] = *v
    }
	
    return result
}