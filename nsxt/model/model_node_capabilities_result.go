/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// List of capabilities of a fabric node
type NodeCapabilitiesResult struct {
	// Node capability results
	Capabilities []NodeCapability `json:"capabilities"`
}
