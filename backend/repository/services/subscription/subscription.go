/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package subscription

// Librerary import
import (
	"backend/repository/config"
	"backend/repository/services/utils"
)

// Vars
var (	
	_insertSubscription = config.ConfigMapSql()["insert_subscription"]
	_updateSubscription = config.ConfigMapSql()["update_subscription"]
)

type Subscription struct {}

/**
  * Get data of all routes
*/
func (p Subscription) InsertSubscription(data ...*interface{}) (map[string]interface{}, error) {
	return utils.GetDataExec(_insertSubscription, utils.ConvertInterfaceDefault(data...)...)
}

/**
  * Get data of all routes
*/
func (p Subscription) UpdateSubscription(data ...*interface{}) (map[string]interface{}, error) {
	return utils.GetDataExec(_updateSubscription, utils.ConvertInterfaceDefault(data...)...)
}