/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type SupportBundleFileTransferAuthenticationScheme struct {
	// User name to authenticate with
	Username string `json:"username"`
	// Authentication scheme name
	SchemeName string `json:"scheme_name"`
	// Password to authenticate with
	Password string `json:"password,omitempty"`
}
