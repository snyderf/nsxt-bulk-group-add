/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MacTableEntry struct {
	// The virtual tunnel endpoint MAC address
	VtepMacAddress string `json:"vtep_mac_address,omitempty"`
	// The virtual tunnel endpoint IP address
	VtepIp string `json:"vtep_ip,omitempty"`
	// The MAC address
	MacAddress string `json:"mac_address"`
}
