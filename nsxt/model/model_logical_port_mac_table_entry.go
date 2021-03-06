/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LogicalPortMacTableEntry struct {
	// The type of the MAC address
	MacType string `json:"mac_type"`
	// The MAC address
	MacAddress string `json:"mac_address"`
}
