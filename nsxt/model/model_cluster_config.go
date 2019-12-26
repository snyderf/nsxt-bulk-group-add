/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ClusterConfig struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// True if control cluster nodes may be added or removed
	ControlClusterChangesAllowed bool `json:"control_cluster_changes_allowed,omitempty"`
	// Configuration of each node in cluster
	Nodes []ClusterNodeInfo `json:"nodes,omitempty"`
	// True if management cluster nodes may be added or removed
	MgmtClusterChangesAllowed bool `json:"mgmt_cluster_changes_allowed,omitempty"`
	// Unique identifier of this cluster
	ClusterId string `json:"cluster_id,omitempty"`
}
