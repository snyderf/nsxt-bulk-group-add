/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type HeatMapTransportZoneStatus struct {
	// Number of transport nodes that are degraded
	DegradedCount int32 `json:"degraded_count,omitempty"`
	// Number of transport nodes that are down
	DownCount int32 `json:"down_count,omitempty"`
	// Number of transport nodes with unknown status
	UnknownCount int32 `json:"unknown_count,omitempty"`
	// Number of transport nodes that are up
	UpCount int32 `json:"up_count,omitempty"`
}
