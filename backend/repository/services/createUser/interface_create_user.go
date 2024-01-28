/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package createUser

/**  Interface create user  **/
type InterfaceCreateUser interface {
	InsertNewUser(...*string) (map[string]interface{}, error)
}

/**
  * Insert New User
*/
func InsertNewUser(p InterfaceCreateUser, data ...*string) (map[string]interface{}, error) {
	return p.InsertNewUser(data...)
}