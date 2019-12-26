/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Abstract base type for Weekly or Interval Backup Schedule
type BackupSchedule struct {
	// Schedule type
	ResourceType string `json:"resource_type"`
}
