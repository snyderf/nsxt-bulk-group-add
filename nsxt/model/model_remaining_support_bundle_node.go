/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type RemainingSupportBundleNode struct {
	// Status of node
	Status string `json:"status,omitempty"`
	// Display name of node
	NodeDisplayName string `json:"node_display_name,omitempty"`
	// UUID of node
	NodeId string `json:"node_id,omitempty"`
}
