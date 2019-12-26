/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ServiceInstanceNSGroups contains list of NS Groups referenced in North-South Service Insertion Rules for a particular Service Instance.
type ServiceInstanceNsGroups struct {
	// List of NSGroups Used in ServiceInsertion Rules.
	Nsroups []NsGroupInfo `json:"nsroups,omitempty"`
}
