/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// IP assignment specification for Static IP List.
type StaticIpListSpec struct {
	ResourceType string `json:"resource_type"`
	// Subnet mask
	SubnetMask string `json:"subnet_mask"`
	// List of IPs for transport node host switch virtual tunnel endpoints
	IpList []string `json:"ip_list"`
	// Gateway IP
	DefaultGateway string `json:"default_gateway"`
}
