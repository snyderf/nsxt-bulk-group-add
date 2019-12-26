/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Supported attributes and sub-attributes for NSProfile
type NsSupportedAttributes struct {
	// The type represent pre-defined list of supported attributes and sub-attributes that can be used while creating NSProfile 
	NsAttributes []NsAttributes `json:"ns_attributes"`
}