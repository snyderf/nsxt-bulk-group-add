/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// IPSec VPN Peer Endpoint covers configuration to be applied locally to establish a session with remote endpoint on peer site.
type IpSecVpnPeerEndpoint struct {
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
	// IPSec Pre-shared key. Maximum length of this field is 128 characters.
	Psk string `json:"psk,omitempty"`
	// Peer identifier.
	PeerId string `json:"peer_id"`
	// Tunnel profile id to be used. By default it will point to system default profile.
	IpsecTunnelProfileId string `json:"ipsec_tunnel_profile_id,omitempty"`
	// Authentication mode used for the peer authentication. For PSK (Pre Shared Key) authentication mode, 'psk' property is mandatory and for the CERTIFICATE authentication mode, 'peer_id' property is mandatory.
	AuthenticationMode string `json:"authentication_mode,omitempty"`
	// IPV4 address of peer endpoint on remote site.
	PeerAddress string `json:"peer_address"`
	// Connection initiation mode used by local endpoint to establish ike connection with peer endpoint. INITIATOR - In this mode local endpoint initiates tunnel setup and will also respond to incoming tunnel setup requests from peer gateway. RESPOND_ONLY - In this mode, local endpoint shall only respond to incoming tunnel setup requests. It shall not initiate the tunnel setup. ON_DEMAND - In this mode local endpoint will initiate tunnel creation once first packet matching the policy rule is received and will also respond to incoming initiation request. 
	ConnectionInitiationMode string `json:"connection_initiation_mode,omitempty"`
	// Dead peer detection (DPD) profile id. Default will be set according to system default policy.
	DpdProfileId string `json:"dpd_profile_id,omitempty"`
	// IKE profile id to be used. Default will be set according to system default policy.
	IkeProfileId string `json:"ike_profile_id,omitempty"`
}
