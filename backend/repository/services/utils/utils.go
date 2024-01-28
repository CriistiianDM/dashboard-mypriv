/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package utils

// Librerary import
import (
	"backend/repository/query"
	"database/sql"
	"fmt"
)

var (
	insQuery = query.GeneralQuery{}
)

/**
  * Return data of query
*/
func GetDataQuery(query_ string , args ...interface{}) ([]map[string]interface{}, error) {
	// Execute the query
	result, is_error := query.Query(insQuery, query_, args...)
	columsResut, column_ := instanceColumsResult(result)
	mapArray := make([]map[string]interface{}, 0)

	// If error
	if is_error == nil {
		defer result.Close()
		for result.Next() {
			err := result.Scan(columsResut...)
			if err == nil {
				row := make(map[string]interface{})
				for i, col := range column_ {
					row[col] = *(columsResut[i].(*interface{}))
				}
				mapArray = append(mapArray, row)
			}
			fmt.Println("Error", err)
		}
	}
	
	// Return result
	return mapArray, is_error
}

/**
  * Return status of exec
*/
func GetDataExec(_query string, args ...interface{}) (map[string]interface{} , error) {
	var is_error error
	response := map[string]interface{}{
		"status": false,
	}
	
	//Execute Query
	result, is_error := query.Exec(insQuery, _query , args...)

	// If error
	if is_error == nil {
		response["status"] = true
		fmt.Println(result)
	}

	// Return result
	return response, is_error
}


/**
  * ConvertColumns to []interface{}
*/
func instanceColumsResult(result *sql.Rows) ([]interface{},[]string) {
	var response []interface{}
	columns_, err := result.Columns()

	// If error
	if err == nil {
		response = make([]interface{} , len(columns_))
		for i := range response {
			var val interface{}
			response[i] = &val
		}
	}

	return response , columns_
}