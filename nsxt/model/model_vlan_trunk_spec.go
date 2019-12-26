/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// VlanTrunkspec is used for specifying trunk VLAN id ranges.
type VlanTrunkSpec struct {
	// Trunk VLAN id ranges
	VlanRanges []TrunkVlanRange `json:"vlan_ranges"`
}