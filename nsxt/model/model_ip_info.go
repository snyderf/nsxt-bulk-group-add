/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type IpInfo struct {
	// IPv4 Addresses
	IpAddresses []string `json:"ip_addresses"`
	// Subnet Prefix Length
	PrefixLength int64 `json:"prefix_length"`
}
