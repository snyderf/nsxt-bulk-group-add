/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type RuntimeInterfaceOperationalStatus struct {
	// The Operational status of the interface
	Status string `json:"status"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Index of the interface
	InterfaceIndex int64 `json:"interface_index,omitempty"`
}
