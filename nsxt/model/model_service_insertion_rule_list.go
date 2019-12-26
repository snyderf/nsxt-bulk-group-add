/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// List of ServiceInsertion Rules.
type ServiceInsertionRuleList struct {
	// List of ServiceInsertion rules in the section. Only homogeneous rules are supported.
	Rules []ServiceInsertionRule `json:"rules"`
}
