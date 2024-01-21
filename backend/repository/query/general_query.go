/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package query

import (
	"backend/db"
	_ "reflect"
	"database/sql"
	_ "fmt"
)
var dbInstance = db.GetConnection()

/**
  * Define querys
*/
type GeneralQuery struct {}

/**
  * Return all params in request
*/
func (p GeneralQuery) Query(sqlQuery string, args ...interface{}) (*sql.Rows, error) {
	dbInstance = db.GetConnection()
	return dbInstance.Query(sqlQuery, args...)
}

/**
  * Execute query
*/
func (p GeneralQuery) Exec(sqlQuery string, args ...interface{}) (sql.Result, error) {
	dbInstance = db.GetConnection()
	return dbInstance.Exec(sqlQuery, args...)
}