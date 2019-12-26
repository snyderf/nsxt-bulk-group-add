/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PoolMemberGroup struct {
	// Load balancer pool support grouping object as dynamic pool members. The IP list of the grouping object such as NSGroup would be used as pool member IP setting. 
	GroupingObject *ResourceReference `json:"grouping_object"`
	// Ip revision filter is used to filter IPv4 or IPv6 addresses from the grouping object. If the filter is not specified, both IPv4 and IPv6 addresses would be used as server IPs. The link local and loopback addresses would be always filtered out. 
	IpRevisionFilter string `json:"ip_revision_filter,omitempty"`
	// The size is used to define the maximum number of grouping object IP address list. These IP addresses would be used as pool members. If the grouping object includes more than certain number of IP addresses, the redundant parts would be ignored and those IP addresses would not be treated as pool members. If the size is not specified, one member is budgeted for this dynamic pool so that the pool has at least one member even if some other dynamic pools grow beyond the capacity of load balancer service. Other members are picked according to available dynamic capacity. The unused members would be set to DISABLED so that the load balancer system itself is not overloaded during runtime. 
	MaxIpListSize int64 `json:"max_ip_list_size,omitempty"`
	// If port is specified, all connections will be sent to this port. If unset, the same port the client connected to will be used, it could be overridden by default_pool_member_ports setting in virtual server. The port should not specified for multiple ports case. 
	Port int32 `json:"port,omitempty"`
	// The list is used to show the customized pool member settings. User can only user pool member action API to update the admin state for a specific IP address. 
	CustomizedMembers []PoolMemberSetting `json:"customized_members,omitempty"`
}