/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package dataUserEmail;

// Librerary import
import (
	"backend/repository/services/utils"
	"backend/repository/config"
)

var (
	_query_data = config.ConfigMapSql()["get_data_email"]
)

/* Class DataUserEmail */
type DataUserEmail struct {}

/**
  * Get all data per email
*/
func (p DataUserEmail) GetDataUserEmail(email string) ([]map[string]interface{}, error) {
	return utils.GetDataQuery(_query_data, email);
}