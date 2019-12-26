/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MigrationUnit struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Info of the group to which this migration unit belongs
	Group *ResourceReference `json:"group,omitempty"`
	// List of warnings indicating issues with the migration unit that may result in migration failure
	Warnings []string `json:"warnings,omitempty"`
	// This is component version e.g. if migration unit is of type HOST, then this is host version.
	CurrentVersion string `json:"current_version,omitempty"`
	// Metadata about migration unit
	Metadata []KeyValuePair `json:"metadata,omitempty"`
	// Migration unit type
	Type_ string `json:"type,omitempty"`
	// Identifier of the migration unit
	Id string `json:"id,omitempty"`
	// Name of the migration unit
	DisplayName string `json:"display_name,omitempty"`
}
