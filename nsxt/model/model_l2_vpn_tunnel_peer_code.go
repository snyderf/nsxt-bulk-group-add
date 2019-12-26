/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// L2Vpn tunnel peer code
type L2VpnTunnelPeerCode struct {
	// Transport tunnel
	TransportTunnel *ResourceReference `json:"transport_tunnel"`
	// Copy this code to paste on the remote end of the tunnel. This is a base64 encoded string which has all the configuration for tunnel. E.g tap device local/peer ips and protocol, encryption algorithm, etc. The peer code also contains a pre-shared key; be careful when sharing or storing it.
	PeerCode string `json:"peer_code"`
}