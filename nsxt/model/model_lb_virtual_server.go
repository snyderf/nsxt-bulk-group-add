/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LbVirtualServer struct {
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
	// Assigned Internet Protocol in IP header, TCP, UDP are supported. 
	IpProtocol string `json:"ip_protocol,omitempty"`
	// Only L7 virtual server could be configured with customized server side TCP profile. 
	ServerTcpProfileId string `json:"server_tcp_profile_id,omitempty"`
	// When load balancer can not select a backend server to serve the request in default pool or pool in rules, the request would be served by sorry server pool. 
	SorryPoolId string `json:"sorry_pool_id,omitempty"`
	// This is a deprecated property, please use 'ports' instead. Port setting could be single port for both L7 mode and L4 mode. For L4 mode, a single port range is also supported. The port setting could be a single port or port range such as \"80\", \"1234-1236\". If port is configured and ports are not specified, both port and ports in response payload would return the same port value. If both port and ports are configured, ports setting would take effect with higher priority. 
	Port string `json:"port,omitempty"`
	// The setting is used when load balancer acts as an SSL client and establishing a connection to the backend server. 
	ServerSslProfileBinding *ServerSslProfileBinding `json:"server_ssl_profile_binding,omitempty"`
	// The server pool(LbPool) contains backend servers. Server pool consists of one or more servers, also referred to as pool members, that are similarly configured and are running the same application. 
	PoolId string `json:"pool_id,omitempty"`
	// Only L7 virtual server could be configured with customized client side TCP profile. 
	ClientTcpProfileId string `json:"client_tcp_profile_id,omitempty"`
	// This is a deprecated property, please use 'default_pool_member_ports' instead. If default_pool_member_port is configured and default_pool_member_ports are not specified, both default_pool_member_port and default_pool_member_ports in response payload would return the same port value. If both are specified, default_pool_member_ports setting would take effect with higher priority. 
	DefaultPoolMemberPort string `json:"default_pool_member_port,omitempty"`
	// Whether access log is enabled
	AccessLogEnabled bool `json:"access_log_enabled,omitempty"`
	// The application profile defines the application protocol characteristics. It is used to influence how load balancing is performed. Currently, LbFastTCPProfile, LbFastUDPProfile and LbHttpProfile, etc are supported. 
	ApplicationProfileId string `json:"application_profile_id"`
	// To ensure one virtual server does not over consume resources, affecting other applications hosted on the same LBS, connections to a virtual server can be capped. If it is not specified, it means that connections are unlimited. 
	MaxConcurrentConnections int64 `json:"max_concurrent_connections,omitempty"`
	// Load balancer rules allow customization of load balancing behavior using match/action rules. Currently, load balancer rules are supported for only layer 7 virtual servers with LbHttpProfile. 
	RuleIds []string `json:"rule_ids,omitempty"`
	// To ensure one virtual server does not over consume resources, connections to a member can be rate limited. If it is not specified, it means that connection rate is unlimited. 
	MaxNewConnectionRate int64 `json:"max_new_connection_rate,omitempty"`
	// Persistence profile is used to allow related client connections to be sent to the same backend server. 
	PersistenceProfileId string `json:"persistence_profile_id,omitempty"`
	// The setting is used when load balancer acts as an SSL server and terminating the client SSL connection 
	ClientSslProfileBinding *ClientSslProfileBinding `json:"client_ssl_profile_binding,omitempty"`
	// If default_pool_member_ports are configured, both default_pool_member_port and default_pool_member_ports in the response payload would include port settings, notice that the value of default_pool_member_port is the first element of default_pool_member_ports. 
	DefaultPoolMemberPorts []string `json:"default_pool_member_ports,omitempty"`
	// virtual server IP address
	IpAddress string `json:"ip_address"`
	// Port setting could be a single port for both L7 mode and L4 mode. For L4 mode, multiple ports or port ranges are also supported such as \"80\", \"443\", \"1234-1236\". If ports is configured, both port and ports in the response payload would include port settings, notice that the port field value is the first element of ports. 
	Ports []string `json:"ports,omitempty"`
	// whether the virtual server is enabled
	Enabled bool `json:"enabled,omitempty"`
}
