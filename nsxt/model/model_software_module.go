/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Software module details
type SoftwareModule struct {
	// Name of the module in the node
	ModuleName string `json:"module_name"`
	// Version of the module in the node
	ModuleVersion string `json:"module_version"`
}