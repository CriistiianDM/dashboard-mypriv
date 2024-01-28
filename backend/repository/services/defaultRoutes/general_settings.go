/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package defaultRoutes

// Librerary import
import (
	"backend/repository/config"
	"backend/repository/services/utils"
)

// Vars
var (
	_get_routes = config.ConfigMapSql()["default_routes"]
)

/** Import Settings of db */
type GeneralRoutes struct {}

/**
  * Get data of all routes
*/
func (p GeneralRoutes) GetAllRoutes() ([]map[string]interface{}, error) {
	return utils.GetDataQuery(_get_routes)
}