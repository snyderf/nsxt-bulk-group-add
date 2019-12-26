/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// A set of operations to be performed in a single batch
type BatchRequest struct {
	Requests []BatchRequestItem `json:"requests,omitempty"`
	// Flag to decide if we will continue processing subsequent requests in case of current error for atomic = false.
	ContinueOnError bool `json:"continue_on_error,omitempty"`
}