/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Allocation rule on edge cluster which will be considered in auto placement of TIER1 logical routers, DHCP and MDProxy. 
type AllocationRule struct {
	// Set action for each allocation rule
	Action *AllocationRuleAction `json:"action"`
}
