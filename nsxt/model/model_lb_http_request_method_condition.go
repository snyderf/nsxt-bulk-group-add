/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// This condition is used to match method of HTTP requests. If the method of an HTTP request is same as the method specified in this condition, the HTTP request match this condition. For example, if the method field is set to GET in this condition, any HTTP request with GET method matches the condition. 
type LbHttpRequestMethodCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// Type of HTTP request method
	Method string `json:"method"`
}
