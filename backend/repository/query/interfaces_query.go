/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package query

// Librerary import
import (
	"database/sql"
)

/**
  * Define the methods of General Query
*/
type InterfaceQuery interface {
	Query(sqlQuery string , args ...interface{}) (*sql.Rows, error)
	Exec(sqlQuery string , args ...interface{}) (sql.Result, error)
}

/**
  * Return query
*/
func Query(p InterfaceQuery, sqlQuery string , args ...interface{}) (*sql.Rows, error) {
	return p.Query(sqlQuery, args...)
}

/**
  * Return exec
*/
func Exec(p InterfaceQuery, sqlQuery string , args ...interface{}) (sql.Result, error) {
	return p.Exec(sqlQuery, args...)
}