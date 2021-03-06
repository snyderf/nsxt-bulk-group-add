/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type RedistributionRule struct {
	// Array of redistribution protocols
	Sources []string `json:"sources"`
	// RouteMap Id for the filter
	RouteMapId string `json:"route_map_id,omitempty"`
	// Description
	Description string `json:"description,omitempty"`
	// Display name
	DisplayName string `json:"display_name,omitempty"`
	// Destination redistribution protocol
	Destination string `json:"destination"`
	// Address family for Route Redistribution
	AddressFamily string `json:"address_family,omitempty"`
}
