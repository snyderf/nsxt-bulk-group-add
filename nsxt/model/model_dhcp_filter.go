/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// DHCP filtering configuration
type DhcpFilter struct {
	// Indicates whether DHCP client blocking is enabled
	ClientBlockEnabled bool `json:"client_block_enabled"`
	// Indiactes whether DHCP v6 client blocking is enabled
	V6ClientBlockEnabled bool `json:"v6_client_block_enabled,omitempty"`
	// Indicates whether DHCP server blocking is enabled
	ServerBlockEnabled bool `json:"server_block_enabled"`
	// Indiactes whether DHCP V6 server blocking is enabled
	V6ServerBlockEnabled bool `json:"v6_server_block_enabled,omitempty"`
}