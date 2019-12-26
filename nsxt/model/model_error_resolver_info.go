/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Metadata related to a given error_id
type ErrorResolverInfo struct {
	// The error id for which metadata information is needed
	ErrorId int64 `json:"error_id"`
	// Indicates whether there is a resolver associated with the error or not
	ResolverPresent bool `json:"resolver_present"`
	// User supplied metadata that might be required by the resolver
	UserMetadata *ErrorResolverUserMetadata `json:"user_metadata,omitempty"`
}
