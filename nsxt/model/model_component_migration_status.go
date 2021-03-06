/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ComponentMigrationStatus struct {
	// Migration status of component
	Status string `json:"status,omitempty"`
	// Indicator of migration progress in percentage
	PercentComplete float32 `json:"percent_complete,omitempty"`
	// Can the migration of the remaining units in this component be skipped
	CanSkip bool `json:"can_skip,omitempty"`
	// Details about the migration status
	Details string `json:"details,omitempty"`
	// Component type for the migration status
	ComponentType string `json:"component_type,omitempty"`
}
