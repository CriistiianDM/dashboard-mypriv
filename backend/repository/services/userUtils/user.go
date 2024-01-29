/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package userUtils

// Librerary import
import (
	"backend/repository/config"
	"backend/repository/services/utils"
)

var (
	_query_follower = config.ConfigMapSql()["get_follower_user"]
	_query_subscriptions = config.ConfigMapSql()["get_subscriptions_user"]
	_query_dashboard = config.ConfigMapSql()["get_dashboard_user"]
	_query_email = config.ConfigMapSql()["get_data_email"]
)

/* Class DataUserEmail */
type UserGeneralUtils struct {}

/**
  * Get Follower per user_account
*/
func (p UserGeneralUtils) GetDataFollwer(user_account int) ([]map[string]interface{}, error) {
	return utils.GetDataQuery(_query_follower, user_account);
}

/**
  * Get Subscription per user_account
*/
func (p UserGeneralUtils) GetDataSubscription(user_account int) ([]map[string]interface{}, error) {
	return utils.GetDataQuery(_query_subscriptions, user_account);
}

/**
  * Get Dashboard per user_account
*/
func (p UserGeneralUtils) GetDataDashboard(user_account int) ([]map[string]interface{}, error) {
	return utils.GetDataQuery(_query_dashboard, user_account);
}

/**
  * Get all data per email
*/
func (p UserGeneralUtils) GetDataUserEmail(email string) ([]map[string]interface{}, error) {
	return utils.GetDataQuery(_query_email, email);
}