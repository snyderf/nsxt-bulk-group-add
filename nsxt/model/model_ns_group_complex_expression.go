/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Complex expressions to represent NSGroup membership
type NsGroupComplexExpression struct {
	ResourceType string `json:"resource_type"`
	// Represents expressions which are to be logically 'AND'ed.The array cannot contain NSGroupComplexExpression.Only NSGroupTagExpression and NSGroupSimpleExpressions are accepted. 
	Expressions []NsGroupExpression `json:"expressions"`
}