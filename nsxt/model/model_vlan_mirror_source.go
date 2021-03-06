/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VlanMirrorSource struct {
	// Resource types of mirror source
	ResourceType string `json:"resource_type"`
	// Source VLAN ID list
	VlanIds []int64 `json:"vlan_ids"`
}
