/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Port Connection Entities (to help draw a visual picture of entities between two ports)
type PortConnectionEntities struct {
	Containers *PortConnectionContainersEntities `json:"containers"`
	Hypervisors []PortConnectionHypervisor `json:"hypervisors"`
	Errors []PortConnectionError `json:"errors"`
	LogicalSwitches []PortConnectionLogicalSwitch `json:"logical_switches"`
	EdgeNodeGroups []PortConnectionEdgeNodeGroup `json:"edge_node_groups,omitempty"`
	Routers []PortConnectionRouter `json:"routers,omitempty"`
	Vms []VirtualMachine `json:"vms"`
	Tunnels []PortConnectionTunnel `json:"tunnels"`
	PhysicalHosts *PortConnectionBmEntities `json:"physical_hosts,omitempty"`
}