/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type BasicAuthenticationScheme struct {
	// Authentication scheme name
	SchemeName string `json:"scheme_name"`
	// User name to authenticate with
	Username string `json:"username"`
	// Password to authenticate with
	Password string `json:"password,omitempty"`
}
