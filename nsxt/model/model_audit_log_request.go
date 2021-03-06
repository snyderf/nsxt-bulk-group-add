/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AuditLogRequest struct {
	// Audit logs should meet the filter condition
	LogFilter string `json:"log_filter,omitempty"`
	// Include logs with timstamps not past the age limit in days
	LogAgeLimit int64 `json:"log_age_limit,omitempty"`
	// Type of log filter
	LogFilterType string `json:"log_filter_type,omitempty"`
}
