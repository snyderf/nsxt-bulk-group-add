/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Protocol on which a particular ServiceInsertion Rule should apply to.
type ServiceInsertionService struct {
	// Display name of the NSX resource.
	TargetDisplayName string `json:"target_display_name,omitempty"`
	// Will be set to false if the referenced NSX resource has been deleted.
	IsValid bool `json:"is_valid,omitempty"`
	// Identifier of the NSX resource.
	TargetId string `json:"target_id,omitempty"`
	// Type of the NSX resource.
	TargetType string `json:"target_type,omitempty"`
	// Distributed Service API accepts raw protocol and ports as part of NS service element in Distributed Service Rule that describes traffic corresponding to an NSService. 
	Service *NsServiceElement `json:"service,omitempty"`
}
