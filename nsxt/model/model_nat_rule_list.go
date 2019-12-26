/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type NatRuleList struct {
	// Add new NatRules to the list in Bulk creation. 
	Rules []NatRule `json:"rules"`
}
