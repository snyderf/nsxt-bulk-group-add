/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Transport tunnel status.
type L2VpnTransportTunnelStatus struct {
	// Resource types of L2VPN Transport tunnels
	ResourceType string `json:"resource_type"`
	// Transport tunnel id.
	TunnelId *ResourceReference `json:"tunnel_id,omitempty"`
}