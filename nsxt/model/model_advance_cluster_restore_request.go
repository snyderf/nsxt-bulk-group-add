/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdvanceClusterRestoreRequest struct {
	// List of instructions and their associated data
	Data []AdvanceClusterRestoreInput `json:"data"`
}
