/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LogicalRouterPortStatisticsSummary struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Rx *LogicalRouterPortCounters `json:"rx,omitempty"`
	Tx *LogicalRouterPortCounters `json:"tx,omitempty"`
	// The ID of the logical router port
	LogicalRouterPortId string `json:"logical_router_port_id"`
}
