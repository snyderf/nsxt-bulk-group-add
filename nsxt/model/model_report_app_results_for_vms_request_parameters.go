/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Parameters to query results of an application discovery session. It has a list of vm ids. 
type ReportAppResultsForVmsRequestParameters struct {
	// Vm external Ids
	VmIds []string `json:"vm_ids,omitempty"`
}