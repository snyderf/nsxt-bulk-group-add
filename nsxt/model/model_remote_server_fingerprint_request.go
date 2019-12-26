/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Remote server
type RemoteServerFingerprintRequest struct {
	// Server port
	Port int64 `json:"port,omitempty"`
	// Remote server hostname or IP address
	Server string `json:"server"`
}