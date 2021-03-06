/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PackageLoggingLevels struct {
	// Logging levels per package
	LoggingLevel string `json:"logging_level,omitempty"`
	// Package name
	PackageName string `json:"package_name,omitempty"`
}
