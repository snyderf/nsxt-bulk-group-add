/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Contains the details of resource field 
type PolicyFineTuningResourceInfoDetail struct {
	// It will represent resource with name and fields. 
	FieldName string `json:"field_name"`
	// List of all field of any resource
	SubType *PolicyFineTuningResourceInfo `json:"sub_type"`
}
