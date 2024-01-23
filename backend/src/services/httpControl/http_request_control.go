/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/
package httpControl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

/**
 *
 * General Control for HTTP Request
 */
type HttpRequestControl struct {
	BodyRequest *gin.Context
}

/**
 * Return all params in request
 * @return map[string]string
 */
func (p HttpRequestControl) GetParamsRequest() map[string]string {
	_response := make(map[string]string)
	// Validate if body request is empty
	if p.BodyRequest != nil {
		// Get body request
		c := p.BodyRequest
		params := c.Params

		// Create a map with the parameters
		for _, p := range params {
			_response[p.Key] = p.Value
		}
	}
	return _response
}

/**
 * Return Body of request
 * @return map[string]interface{}
 */
func (p HttpRequestControl) GetBodyRequest() map[string]interface{} {
	// Initialize body request
	_body := make(map[string]interface{})
	// Validate if body request is empty
	if p.BodyRequest != nil {

		// Get body request
		c := p.BodyRequest

		// Get body request
		body, err := c.GetRawData()

		// Validate if body request is not empty
		if err == nil {
						
			err := json.Unmarshal(body, &_body)
			if err != nil {
				_body = make(map[string]interface{})
			}
		} else {
			fmt.Println("No can get body request", err)
		}

	}
	return _body
}
