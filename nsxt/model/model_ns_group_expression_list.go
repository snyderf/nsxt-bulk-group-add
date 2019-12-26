/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// List of NSGroupExpressions
type NsGroupExpressionList struct {
	// List of NSGroupExpressions to be passed to add and remove APIs 
	Members []NsGroupExpression `json:"members"`
}
