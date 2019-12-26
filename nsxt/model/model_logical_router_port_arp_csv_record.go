/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LogicalRouterPortArpCsvRecord struct {
	// The IP address
	Ip string `json:"ip"`
	// The MAC address
	MacAddress string `json:"mac_address"`
}