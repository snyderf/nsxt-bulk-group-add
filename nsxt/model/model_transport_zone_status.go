/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Transport zone runtime status information
type TransportZoneStatus struct {
	// Count of logical ports in the transport zone
	NumLogicalPorts int32 `json:"num_logical_ports,omitempty"`
	// Unique ID identifying the transport zone
	TransportZoneId string `json:"transport_zone_id,omitempty"`
	// Count of logical switches in the transport zone
	NumLogicalSwitches int32 `json:"num_logical_switches,omitempty"`
	// Count of transport nodes in the transport zone
	NumTransportNodes int32 `json:"num_transport_nodes,omitempty"`
}