/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package followers	

/** Interface Followers **/
type InterfaceFollower interface {
	InsertFollower(...*int) (map[string]interface{}, error)
	UpdateFollower(...*interface{}) (map[string]interface{}, error) 
}

/**
  * Get data of all routes
*/
func InsertFollower(p InterfaceFollower , data ...*int) (map[string]interface{}, error) {
	return p.InsertFollower(data...)
}

/**
  * Get data of all routes
*/
func UpdateFollower(p InterfaceFollower , data ...*interface{}) (map[string]interface{}, error) {
	return p.UpdateFollower(data...)
}