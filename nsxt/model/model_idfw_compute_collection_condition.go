/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Status of the Identity Firewall enabled Compute collection.
type IdfwComputeCollectionCondition struct {
	// IDFW enabled Compute collection status.
	Status string `json:"status"`
	// Status of the Compute collection.
	StatusDetail string `json:"status_detail,omitempty"`
}
