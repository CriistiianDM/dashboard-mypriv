/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package defaultRoutes

// Librerary import
import (
	"backend/repository/query"
	"backend/repository/config"
	"fmt"
)

// Vars
var (
	instanceQuery = query.GeneralQuery{}
	_get_routes = config.ConfigMapSql()["default_routes"]
)

/** Import Settings of db */
type GeneralRoutes struct {
	route string
}

/**
  * Get data of all routes
*/
func (p GeneralRoutes) GetAllRoutes() (map[string]interface{}, error) {
	// Execute the query
	response := make(map[string]interface{})
	result, is_error := query.Query(instanceQuery, _get_routes)

	// If error
	if is_error == nil {
		defer result.Close()
		for result.Next() {
			err := result.Scan(&p.route)
			if err != nil {
				fmt.Println("Error", err)
			}
			response[p.route] = p
		}
	}
	// Return result
	return response, is_error
}