/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Attributes/sub-attributes data holder structure for NSProfile
type NsAttributesData struct {
	// Data type of attribute/sub attribute key
	Datatype string `json:"datatype"`
	// Multiple attribute/sub attribute values can be specified as elements of array. 
	Value []string `json:"value"`
	// NSProfile attribute/sub attribute keys. 
	Key string `json:"key"`
	// Description for NSProfile attributes
	Description string `json:"description,omitempty"`
}
