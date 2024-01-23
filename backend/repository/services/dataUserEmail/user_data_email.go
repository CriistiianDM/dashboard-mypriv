/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package dataUserEmail;

// Librerary import
import (
	"backend/repository/query"
	"backend/repository/config"
	"fmt"
)

var (
	_query_data_email = config.ConfigMapSql()["get_data_email"]
	instanceQuery = query.GeneralQuery{}
)

/* Class DataUserEmail */
type DataUserEmail struct {
	id int,
    nickname string,
    created_at string,
    first_name string,
    last_name string,
    profile_picture string
    data string
}


/**
  * Get all data per email
*/
func (p DataUserEmail) GetDataUserEmail(email string) (map[string]interface{}, error) {
	// Execute the query
	response := make(map[string]interface{})
	result, is_error := query.Query(instanceQuery, _query_data_email)

	// If error
	if is_error == nil {
		defer result.Close()
		for result.Next() {
			err := result.Scan(&p)
			if err != nil {
				fmt.Println("Error", err)
			}
			response[p.nickname] = p
		}
	}
	// Return result
	return response, is_error
}
