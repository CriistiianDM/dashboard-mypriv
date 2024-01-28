/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package defaultRoutes

/** Interface Setting */
type InterfaceGeneralRoutes interface {
	GetAllRoutes() ([]map[string]interface{}, error)
}

/**
  * Return all params in request
*/
func GetAllRoutes(p InterfaceGeneralRoutes) ([]map[string]interface{}, error) {
	return p.GetAllRoutes()
}