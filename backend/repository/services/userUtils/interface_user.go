/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package userUtils

/** Interface Setting */
type InterfaceUserData interface {
	GetDataFollwer(int) ([]map[string]interface{}, error)
	GetDataSubscription(int) ([]map[string]interface{}, error)
	GetDataDashboard(int) ([]map[string]interface{}, error)
	GetDataUserEmail(string) ([]map[string]interface{}, error)
	GetAllUserData() ([]map[string]interface{}, error)
}

/**
  * Return followers of user per user_account
*/
func GetDataFollwer(p InterfaceUserData, user_account int) ([]map[string]interface{}, error) {
	return p.GetDataFollwer(user_account)
}

/**
  * Return Subscription of user per user_account
*/
func GetDataSubscription(p InterfaceUserData, user_account int) ([]map[string]interface{}, error) {
	return p.GetDataSubscription(user_account)
}

/**
  * Return Subscription of user per user_account
*/
func GetDataDashboard(p InterfaceUserData, user_account int) ([]map[string]interface{}, error) {
	return p.GetDataDashboard(user_account)
}

/**
  * Return data of user per email
*/
func GetDataUserEmail(p InterfaceUserData, email string) ([]map[string]interface{}, error) {
	return p.GetDataUserEmail(email)
}

/**
  * Get all data user
*/
func GetAllUserData(p InterfaceUserData) ([]map[string]interface{}, error) {
	return p.GetAllUserData()
}