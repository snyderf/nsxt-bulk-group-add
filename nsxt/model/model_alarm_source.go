/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Source of the Alarm
type AlarmSource struct {
	// VPN session/tunnel etc. local Ip(v4 or v6) address
	LocalIp string `json:"local_ip,omitempty"`
	// Identifier of the Alarm source entity for e.g. Vpn session id, Vpn tunnel id etc.
	Id string `json:"id,omitempty"`
	// Filesystem mount name
	Mount string `json:"mount,omitempty"`
	// Ip address
	IpAddress string `json:"ip_address,omitempty"`
	// Ip address type v4, v6 etc.
	IpAddressType string `json:"ip_address_type,omitempty"`
	// VPN session/tunnel etc. peer Ip(v4 or v6) address
	PeerIp string `json:"peer_ip,omitempty"`
}