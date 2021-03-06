/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Tunnel information between two given transport nodes
type PortConnectionTunnel struct {
	// Id of the source transport node
	SrcNodeId string `json:"src_node_id"`
	// Tunnel properties between the source and the destination transport node
	TunnelProperties *TunnelProperties `json:"tunnel_properties"`
}
