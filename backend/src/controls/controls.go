/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2024
   * @license: GPL-3.0
*/

package controls;

import (
	"github.com/gin-gonic/gin"
	_ "fmt"
	_ "backend/src/services/httpControl"
)

/**
 * Return Data per email
*/
func GetUserByEmail(c *gin.Context) {
	c.JSON(200, gin.H{"response": true , "message": "Product inserted"})
}