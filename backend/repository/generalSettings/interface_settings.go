/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package generalSettings

/** Interface Setting */
type InterfaceGeneralSettings interface {
	GetAllRoutes() (map[string]interface{}, error)
}

/**
  * Return all params in request
*/
func GetAllRoutes(p InterfaceGeneralSettings) (map[string]interface{}, error) {
	return p.GetAllRoutes()
}