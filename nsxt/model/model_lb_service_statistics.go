/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LbServiceStatistics struct {
	// Statistics of load balancer pools
	Pools []LbPoolStatistics `json:"pools,omitempty"`
	// load balancer service identifier
	ServiceId string `json:"service_id"`
	// Statistics of load balancer virtual servers
	VirtualServers []LbVirtualServerStatistics `json:"virtual_servers,omitempty"`
	// Timestamp when the data was last updated
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Load balancer service statistics counter
	Statistics *LbServiceStatisticsCounter `json:"statistics,omitempty"`
}
