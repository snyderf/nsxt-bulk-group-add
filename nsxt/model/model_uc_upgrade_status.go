/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Upgrade status of upgrade-coordinator
type UcUpgradeStatus struct {
	// Status of UC upgrade
	Status string `json:"status,omitempty"`
	// Current state of UC upgrade
	State string `json:"state,omitempty"`
}
