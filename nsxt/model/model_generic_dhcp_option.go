/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Define DHCP options other than option 121.
type GenericDhcpOption struct {
	// Code of the dhcp option.
	Code int64 `json:"code"`
	// Value of the option.
	Values []string `json:"values"`
}