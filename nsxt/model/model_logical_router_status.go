/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LogicalRouterStatus struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The id of the logical router
	LogicalRouterId string `json:"logical_router_id"`
	// Per Node Status
	PerNodeStatus []LogicalRouterStatusPerNode `json:"per_node_status,omitempty"`
}
