/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ManagementClusterStatus struct {
	// The current status of the management cluster
	Status string `json:"status,omitempty"`
	// Current missing management plane nodes
	OfflineNodes []ManagementPlaneBaseNodeInfo `json:"offline_nodes,omitempty"`
	// The details of the cluster nodes required for cluster initialization
	RequiredMembersForInitialization []ClusterInitializationNodeInfo `json:"required_members_for_initialization,omitempty"`
	// Current alive management plane nodes
	OnlineNodes []ManagementPlaneBaseNodeInfo `json:"online_nodes,omitempty"`
}