/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type NatStatisticsPerLogicalRouter struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Detailed per node statistics
	PerTransportNodeStatistics []NatStatisticsPerTransportNode `json:"per_transport_node_statistics,omitempty"`
	// Rolled-up statistics for all rules on the logical router across all the nodes
	StatisticsAcrossAllNodes *NatCounters `json:"statistics_across_all_nodes,omitempty"`
	// Id for the logical router
	LogicalRouterId string `json:"logical_router_id,omitempty"`
}
