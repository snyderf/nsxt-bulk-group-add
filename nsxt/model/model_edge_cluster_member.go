/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type EdgeClusterMember struct {
	// System generated index for cluster member
	MemberIndex int32 `json:"member_index,omitempty"`
	// display name of edge cluster member
	DisplayName string `json:"display_name,omitempty"`
	// description of edge cluster member
	Description string `json:"description,omitempty"`
	// Identifier of the transport node backed by an Edge node
	TransportNodeId string `json:"transport_node_id"`
}