/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package followers

// Librerary import
import (
	"backend/repository/config"
	"backend/repository/services/utils"
	"fmt"
)

// Vars
var (
	_insertFollower = config.ConfigMapSql()["insert_follower"]
)

type FollersGeneral struct {}

/**
  * Get data of all routes
*/
func (p FollersGeneral) InsertFollower(data ...*int) (map[string]interface{}, error) {
	return utils.GetDataExec(_insertFollower, p.convertInterfaceSlice(data)...)
}

/**
  * Convert []*int to []interface{}
*/
func (p FollersGeneral) convertInterfaceSlice(data []*int) []interface{} {
    result := make([]interface{}, len(data))
    for i, v := range data {
        result[i] = *v
    }
	fmt.Println("aaa", result)
    return result
}