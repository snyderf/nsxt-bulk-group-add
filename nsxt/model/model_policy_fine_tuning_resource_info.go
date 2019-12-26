/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// It represent the resource with details of name and fields it owns. 
type PolicyFineTuningResourceInfo struct {
	// List of all field of any resource
	Fields []PolicyFineTuningResourceInfoDetail `json:"fields"`
	// It will represent resource with name and fields. 
	ResourceName string `json:"resource_name"`
}
