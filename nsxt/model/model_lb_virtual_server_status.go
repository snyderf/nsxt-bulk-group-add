/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LbVirtualServerStatus struct {
	// UP means that all primary members in default pool are in UP status. For L7 virtual server, if there is no default pool, the virtual server would be treated as UP. PARTIALLY_UP means that some(not all) primary members in default pool are in UP status. The size of these active primary members should be larger than or equal to the certain number(min_active_members) which is defined in LbPool. When there are no backup members which are in the UP status, the number(min_active_members) would be ignored. PRIMARY_DOWN means that less than certain(min_active_members) primary members in default pool are in UP status but backup members are in UP status, the connections would be dispatched to backup members. DOWN means that all primary and backup members are in DOWN status. DETACHED means that the virtual server is not bound to any service. DISABLED means that the admin state of the virtual server is disabled. UNKNOWN means that no status reported from transport-nodes. The associated load balancer service may be working(or not working). 
	Status string `json:"status,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// load balancer virtual server identifier
	VirtualServerId string `json:"virtual_server_id"`
}
