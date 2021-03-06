/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type SupportBundleResult struct {
	// Request properties
	RequestProperties *SupportBundleRequest `json:"request_properties,omitempty"`
	// Nodes where bundles were not generated or not copied to remote server
	FailedNodes []FailedNodeSupportBundleResult `json:"failed_nodes,omitempty"`
	// Nodes whose bundles were successfully copied to remote file server
	SuccessNodes []SuccessNodeSupportBundleResult `json:"success_nodes,omitempty"`
	// Nodes where bundle generation is pending or in progress
	RemainingNodes []RemainingSupportBundleNode `json:"remaining_nodes,omitempty"`
}
