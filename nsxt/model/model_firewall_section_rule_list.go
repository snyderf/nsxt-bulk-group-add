/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type FirewallSectionRuleList struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Stateful or Stateless nature of distributed service section is enforced on all rules inside the section. Layer3 sections can be stateful or stateless. Layer2 sections can only be stateless.
	Stateful bool `json:"stateful"`
	// It is a boolean flag which reflects whether a distributed service section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section.
	IsDefault bool `json:"is_default,omitempty"`
	// List of objects where the rules in this section will be enforced. This will take precedence over rule level appliedTo.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Number of rules in this section.
	RuleCount int64 `json:"rule_count,omitempty"`
	// Type of the rules which a section can contain. Only homogeneous sections are supported.
	SectionType string `json:"section_type"`
	// Priority of current section with respect to other sections. In case the field is empty, the list section api should be used to get section priority.
	Priority int64 `json:"priority,omitempty"`
	// This attribute represents enforcement point of firewall section. For example, firewall section enforced on logical port with attachment type bridge endpoint will have 'BRIDGEENDPOINT' value, firewall section enforced on logical router will have 'LOGICALROUTER' value and rest have 'VIF' value.
	EnforcedOn string `json:"enforced_on,omitempty"`
	// Section is locked/unlocked.
	Locked bool `json:"locked,omitempty"`
	// If TCP strict is enabled on a section and a packet matches rule in it, the following check will be performed. If the packet does not belong to an existing session, the kernel will check to see if the SYN flag of the packet is set. If it is not, then it will drop the packet.
	TcpStrict bool `json:"tcp_strict,omitempty"`
	// ID of the user who last modified the lock for the section.
	LockModifiedBy string `json:"lock_modified_by,omitempty"`
	// Section locked/unlocked time in epoch milliseconds.
	LockModifiedTime int64 `json:"lock_modified_time,omitempty"`
	// Comments for section lock/unlock.
	Comments string `json:"comments,omitempty"`
	// This flag indicates whether it is an auto-plumbed section that is associated to a LogicalRouter. Auto-plumbed sections are system owned and cannot be updated via the API.
	Autoplumbed bool `json:"autoplumbed,omitempty"`
	// List of firewall rules in the section. Only homogenous rules are supported.
	Rules []FirewallRule `json:"rules"`
}
