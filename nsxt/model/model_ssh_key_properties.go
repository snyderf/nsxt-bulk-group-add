/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type SshKeyProperties struct {
	// Current password for user (required for users root and admin)
	Password string `json:"password,omitempty"`
	// SSH key label (used to identify the key)
	Label string `json:"label"`
	// SSH key type
	Type_ string `json:"type"`
	// SSH key value
	Value string `json:"value"`
}
