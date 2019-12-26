/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ControlClusterNodeStatus struct {
	// Status of this node's management plane connection
	MgmtConnectionStatus *MgmtConnStatus `json:"mgmt_connection_status,omitempty"`
	// Status of this node's connection to the control cluster
	ControlClusterStatus string `json:"control_cluster_status,omitempty"`
}
