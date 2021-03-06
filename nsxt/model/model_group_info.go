/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// GroupInfo contains information about a particular Group used in Redirection Rules. It also contains information about policy path, if the group is created from Policy.
type GroupInfo struct {
	// Group Data.
	Group *ResourceReference `json:"group,omitempty"`
	// Policy path of a particular Group.
	GroupPolicyPath string `json:"group_policy_path,omitempty"`
}
