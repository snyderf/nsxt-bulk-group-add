/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type FirewallSectionsSummary struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Total number of sections for the section type.
	SectionCount int64 `json:"section_count,omitempty"`
	// Total number of rules in the section.
	RuleCount int64 `json:"rule_count,omitempty"`
	// Type of rules which a section can contain.
	SectionType string `json:"section_type,omitempty"`
}
