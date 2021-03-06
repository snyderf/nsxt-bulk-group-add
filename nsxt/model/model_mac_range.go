/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// A range of MAC addresses with a start and end value
type MacRange struct {
	// Start value for MAC address range
	Start string `json:"start"`
	// End value for MAC address range
	End string `json:"end"`
}
