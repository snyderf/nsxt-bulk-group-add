/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// NSX-Intelligence configuration that can be applied to host nodes. 
type IntelligenceHostConfigurationInfo struct {
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
	// Interval in minute of reporting VM guest context data to NSX-Intelligence. 
	ContextDataCollectionInterval int64 `json:"context_data_collection_interval,omitempty"`
	// A truststore to establish the trust between NSX and NSX-Intelligence brokers. 
	BrokerTruststore string `json:"broker_truststore,omitempty"`
	// A broker certificate to verify the identity of brokers. 
	BrokerCertificate string `json:"broker_certificate,omitempty"`
	// List of linux user uid to collect context data. Empty implies all users. 
	ContextUserUids []string `json:"context_user_uids,omitempty"`
	// List of hashes of processes to collect context data. Empty implies all processes. 
	ContextProcessHashes []string `json:"context_process_hashes,omitempty"`
	// Enable NSX-Intelligence data collection in host nodes. 
	EnableDataCollection bool `json:"enable_data_collection,omitempty"`
	// List of processes to collect context data. Empty implies all processes. 
	ContextProcessNames []string `json:"context_process_names,omitempty"`
	// List of private IP prefix that NSX-Intelligence network flow is collected from. 
	PrivateIpPrefix []IntelligenceFlowPrivateIpPrefixInfo `json:"private_ip_prefix,omitempty"`
	// List of NSX-Intelligence broker endpoints that host nodes contact initially. 
	BrokerBootstrapServers []IntelligenceBrokerEndpointInfo `json:"broker_bootstrap_servers,omitempty"`
	// Maximum inactive network flow to collect in collection interval. 
	MaxInactiveFlowCount int64 `json:"max_inactive_flow_count,omitempty"`
	// List of windows user sid to collect context data. Empty implies all users. 
	ContextUserSids []string `json:"context_user_sids,omitempty"`
	// Interval in minute of reporting network flow data to NSX-Intelligence. 
	FlowDataCollectionInterval int64 `json:"flow_data_collection_interval,omitempty"`
	// Maximum active network flow to collect in collection interval. 
	MaxActiveFlowCount int64 `json:"max_active_flow_count,omitempty"`
}
