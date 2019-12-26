/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Simple expressions to represent NSGroup membership
type NsGroupSimpleExpression struct {
	ResourceType string `json:"resource_type"`
	
	// Reference of the target. Will be populated when the property is a resource id, the op (operator) is EQUALS and populate_references is set to be true. 
	TargetResource *ResourceReference `json:"target_resource,omitempty"`
	
	// Field of the resource on which this expression is evaluated
	TargetProperty string `json:"target_property"`
	
	// Type of the resource on which this expression is evaluated
	TargetType string `json:"target_type"`
	
	// Value that satisfies this expression
	Value string `json:"value"`
	
	// Operator of the expression
	Op string `json:"op"`
}
