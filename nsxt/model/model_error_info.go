/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Error information
type ErrorInfo struct {
	// Timestamp when the error occurred
	Timestamp int64 `json:"timestamp,omitempty"`
	// Error message
	ErrorMessage string `json:"error_message,omitempty"`
}
