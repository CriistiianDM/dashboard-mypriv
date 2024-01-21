/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package generalSettings

// Librerary import
import (
	"backend/repository/query"
	"fmt"
)

// Vars
var (
	_get_routes = "SELECT * FROM get_default_routes()"
	instanceQuery = query.GeneralQuery{}
)

/** Import Settings of db */
type GeneralSettings struct {
	route string
}

/**
  * Get data of all routes
*/
func (p GeneralSettings) GetAllRoutes() (map[string]interface{}, error) {
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