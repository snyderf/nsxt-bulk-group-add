/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Collection of edge nodes backing a logical router
type PortConnectionEdgeNodeGroup struct {
	// Resource reference with details of the entity
	Resource *ManagedResource `json:"resource,omitempty"`
	// Resource ID is mapped to this. (ID is Generated for Edge node groups, since resource will be null)
	Id string `json:"id,omitempty"`
	EdgeNodes []TransportNode `json:"edge_nodes,omitempty"`
	// Id of the logical router
	LogicalRouterId string `json:"logical_router_id"`
}
