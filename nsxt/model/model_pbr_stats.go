/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PbrStats struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Aggregated number of packets processed by the rule.
	PacketCount int64 `json:"packet_count,omitempty"`
	// Aggregated number of bytes processed by the rule.
	ByteCount int64 `json:"byte_count,omitempty"`
	// Rule Identifier of the PBR rule. This is a globally unique number.
	RuleId string `json:"rule_id,omitempty"`
}
