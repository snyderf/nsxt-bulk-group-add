/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Failure Domain is engineered to be isolated from failures in other failure domains, and to provide inexpensive, low-latency network connectivity to other failure domains in the same region. We support failure domain only on edge transport node. Like you can consider one rack as one failure domain and place active-standby contexts like logical router, DHCP and MDProxy in two different failure domains (racks). So failure of a single rack (FD) does not impact any services and other rack (FD) will continue to handle traffic. Initially system creates one default failure domain. For any edge transport node, if failure domains is not given, it will be mapped to default system generated failure domain. You can't update preferred_active_edge_services flag for system generated default failure domain. It will be unset which means that failure domain can be used for both active and standby allocation. 
type FailureDomain struct {
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
	// Set preference for edge transport node failure domain which will be considered while doing auto placement of logical router, DHCP and MDProxy on edge node. true: For preemptive failover mode, active edge cluster member       allocation preferes this failure domain. false: For preemptive failover mode, standby edge cluster member       allocation preferes this failure domain. Default will be unset. It means no explicit preference. 
	PreferredActiveEdgeServices bool `json:"preferred_active_edge_services,omitempty"`
}
