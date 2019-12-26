/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Node log properties
type NodeLogProperties struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Last modified time expressed in milliseconds since epoch
	LastModifiedTime int64 `json:"last_modified_time,omitempty"`
	// Size of log file in bytes
	LogSize int64 `json:"log_size,omitempty"`
	// Name of log file
	LogName string `json:"log_name,omitempty"`
}
