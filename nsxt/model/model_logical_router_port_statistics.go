/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LogicalRouterPortStatistics struct {
	// Per Node Statistics
	PerNodeStatistics []LogicalRouterPortStatisticsPerNode `json:"per_node_statistics,omitempty"`
	// The ID of the logical router port
	LogicalRouterPortId string `json:"logical_router_port_id"`
}
