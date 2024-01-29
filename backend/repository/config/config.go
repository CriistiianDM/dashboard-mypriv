/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package config;

/**
  * Config Map Sql
*/
func ConfigMapSql() map[string]string {
	return map[string]string{
		"default_routes": "SELECT * FROM get_default_routes()",
		"get_data_email": "SELECT * FROM get_user_all_data($1)",
		"create_user": "CALL insert_new_user_account($1,$2,$3,$4,$5,$6)",
		"insert_follower": "CALL insert_row_follow($1,$2)",
	}
}