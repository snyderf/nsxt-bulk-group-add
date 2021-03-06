/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Role
type NsxRole struct {
	// Role name
	Role string `json:"role"`
	// Please use the /user-info/permissions api to get the permission that the user has on each feature.
	Permissions []string `json:"permissions,omitempty"`
}
