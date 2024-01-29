/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package subscription	

/** Interface Followers **/
type InterfaceSubscription interface {
	InsertSubscription(...*interface{}) (map[string]interface{}, error)
	UpdateSubscription(...*interface{}) (map[string]interface{}, error) 
}

/**
  * Get data of all routes
*/
func InsertSubscription(p InterfaceSubscription , data ...*interface{}) (map[string]interface{}, error) {
	return p.InsertSubscription(data...)
}

/**
  * Get data of all routes
*/
func UpdateSubscription(p InterfaceSubscription , data ...*interface{}) (map[string]interface{}, error) {
	return p.UpdateSubscription(data...)
}