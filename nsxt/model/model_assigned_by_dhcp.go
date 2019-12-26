/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// This type can be specified in ip assignment spec of host switch if DHCP based IP assignment is desired for host switch virtual tunnel endpoints.
type AssignedByDhcp struct {
	ResourceType string `json:"resource_type"`
}