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
)

// Vars
var (	
	_insertFollower = config.ConfigMapSql()["insert_follower"]
	_updateFollower = config.ConfigMapSql()["update_follower"]
)

type FollowersGeneral struct {}

/**
  * Insert Followers
*/
func (p FollowersGeneral) InsertFollower(data ...*int) (map[string]interface{}, error) {
	return utils.GetDataExec(_insertFollower, utils.ConvertInterfaceSlice(data)...)
}

/**
  * Update Followers
*/
func (p FollowersGeneral) UpdateFollower(data ...*interface{}) (map[string]interface{}, error) {
	return utils.GetDataExec(_updateFollower, utils.ConvertInterfaceDefault(data...)...)
}