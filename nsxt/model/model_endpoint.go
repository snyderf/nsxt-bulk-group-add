/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// An Endpoint object is part of HostSwitch configuration in TransportNode
type Endpoint struct {
	// Subnet mask
	SubnetMask string `json:"subnet_mask,omitempty"`
	// MAC address
	Mac string `json:"mac,omitempty"`
	// Gateway IP
	DefaultGateway string `json:"default_gateway,omitempty"`
	// Depending upon the EndpointIpConfig used in HostSwitch, IP could be allocated either from DHCP (default) or from Static IP Pool.
	Ip string `json:"ip,omitempty"`
	// Name of the virtual tunnel endpoint
	DeviceName string `json:"device_name,omitempty"`
	// Unique label for this Endpoint
	Label int32 `json:"label,omitempty"`
}
