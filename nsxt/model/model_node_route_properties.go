/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Node network route properties
type NodeRouteProperties struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Source address to prefer when sending to destinations of route
	Src string `json:"src,omitempty"`
	// From address
	FromAddress string `json:"from_address,omitempty"`
	// Routing protocol identifier of route
	Proto string `json:"proto,omitempty"`
	// Route type
	RouteType string `json:"route_type"`
	// Metric value of route
	Metric string `json:"metric,omitempty"`
	// Destination covered by route
	Destination string `json:"destination,omitempty"`
	// Network interface id of route
	InterfaceId string `json:"interface_id,omitempty"`
	// Unique identifier for the route
	RouteId string `json:"route_id,omitempty"`
	// Netmask of destination covered by route
	Netmask string `json:"netmask,omitempty"`
	// Scope of destinations covered by route
	Scope string `json:"scope,omitempty"`
	// Address of next hop
	Gateway string `json:"gateway,omitempty"`
}
