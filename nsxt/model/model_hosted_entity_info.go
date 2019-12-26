/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type HostedEntityInfo struct {
	// Unique identifier of entity
	EntityUuid string `json:"entity_uuid,omitempty"`
	// The type of entity hosted could be MP, CCP, VMC App etc.
	EntityType string `json:"entity_type,omitempty"`
}