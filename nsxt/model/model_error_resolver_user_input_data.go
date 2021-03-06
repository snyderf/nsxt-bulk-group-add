/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Corresponds to one property entered by the user
type ErrorResolverUserInputData struct {
	// The datatype of the given property. Useful for data validation
	DataType string `json:"data_type"`
	// Name of the property supplied by the user
	PropertyName string `json:"property_name"`
	// The value associated with the above property
	PropertyValue string `json:"property_value,omitempty"`
}
