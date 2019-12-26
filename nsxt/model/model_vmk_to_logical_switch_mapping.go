/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Note- transport node templates APIs are deprecated and user is recommended to use transport node profiles APIs instead.
type VmkToLogicalSwitchMapping struct {
	// Only VLAN type of logical switch can be specified here, as migration operation is only supported for these types.
	DestNetworkId string `json:"dest_network_id"`
	// The ID (for example, vmk0) of the vmkernal interface on ESX to migrate.
	SrcInterfaceId string `json:"src_interface_id"`
}
