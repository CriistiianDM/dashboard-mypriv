/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package dataUserEmail

/** Interface Setting */
type InterfaceDataUserEmail interface {
	GetDataUserEmail(string) ([]map[string]interface{}, error)
}

/**
  * Return all params in request
*/
func GetDataUserEmail(p InterfaceDataUserEmail, email string) ([]map[string]interface{}, error) {
	return p.GetDataUserEmail(email)
}