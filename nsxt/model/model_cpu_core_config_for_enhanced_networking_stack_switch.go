/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Non Uniform Memory Access (NUMA) nodes and Logical cpu cores (Lcores) per NUMA node configuration for Enhanced Networking Stack enabled HostSwitch.
type CpuCoreConfigForEnhancedNetworkingStackSwitch struct {
	// Number of Logical cpu cores (Lcores) to be placed on a specified NUMA node
	NumLcores int32 `json:"num_lcores"`
	// Unique index of the Non Uniform Memory Access (NUMA) node
	NumaNodeIndex int32 `json:"numa_node_index"`
}
