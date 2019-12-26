/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Physical NIC specification
type Pnic struct {
	// Uplink name for this Pnic. This name will be used to reference this Pnic in other configurations.
	UplinkName string `json:"uplink_name"`
	// device name or key
	DeviceName string `json:"device_name"`
}
