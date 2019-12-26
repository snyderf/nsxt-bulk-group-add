/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Identity Firewall NSGorup to user mapping to link DirGroup to user session data. 
type IdfwDirGroupUserSessionMapping struct {
	// User ID.
	UserId string `json:"user_id,omitempty"`
	// Directory Group ID.
	DirGroupId string `json:"dir_group_id,omitempty"`
}
