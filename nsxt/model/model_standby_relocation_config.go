/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Standby service contexts relocation setting
type StandbyRelocationConfig struct {
	// The time interval (in minutes) to wait before starting the standby service context relocation process. 
	StandbyRelocationThreshold int64 `json:"standby_relocation_threshold,omitempty"`
}
