/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Pool usage statistics in a pool.
type PoolUsage struct {
	// Total number of IDs in a pool
	TotalIds int64 `json:"total_ids,omitempty"`
	// Total number of allocated IDs in a pool
	AllocatedIds int64 `json:"allocated_ids,omitempty"`
	// Total number of free IDs in a pool
	FreeIds int64 `json:"free_ids,omitempty"`
}
