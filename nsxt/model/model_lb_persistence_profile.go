/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// LB persistence profile contains the information related to load balancer persistence options. Some applications maintain state and require all relevant connections to be sent to the same server as the application state is not synchronized among servers. Persistence can be enabled on a LbVirtualServer by binding a persistence profile to it. LbGenericPersistenceProfile cannot be attached to virtual server directly, it can be only consumed by LB rule action. If a user attaches a generic persistence profile directly to a virtual server, the operation is rejected. 
type LbPersistenceProfile struct {
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
	// The resource_type property identifies persistence profile type. 
	ResourceType string `json:"resource_type"`
	// The persistence shared flag identifies whether the persistence table is shared among virtual-servers referring this profile. If persistence shared flag is not set in the cookie persistence profile bound to a virtual server, it defaults to cookie persistence that is private to each virtual server and is qualified by the pool. This is accomplished by load balancer inserting a cookie with name in the format &lt;name&gt;.&lt;virtual_server_id&gt;.&lt;pool_id&gt;. If persistence shared flag is set in the cookie persistence profile, in cookie insert mode, cookie persistence could be shared across multiple virtual servers that are bound to the same pools. The cookie name would be changed to &lt;name&gt;.&lt;profile-id&gt;.&lt;pool-id&gt;. If persistence shared flag is not set in the sourceIp persistence profile bound to a virtual server, each virtual server that the profile is bound to maintains its own private persistence table. If persistence shared flag is set in the sourceIp persistence profile, all virtual servers the profile is bound to share the same persistence table. If persistence shared flag is not set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using both virtual server ID and profile ID. If persistence shared flag is set in the generic persistence profile, the persistence entries are matched and stored in the table which is identified using profile ID. It means that virtual servers which consume the same profile in the LbRule with this flag enabled are sharing the same persistence table. 
	PersistenceShared bool `json:"persistence_shared,omitempty"`
}
