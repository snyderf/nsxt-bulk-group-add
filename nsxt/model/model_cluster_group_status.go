/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// This type contains the attributes and status of a group.
type ClusterGroupStatus struct {
	// Array of group members and their statuses
	Members []ClusterGroupMemberStatus `json:"members,omitempty"`
	// UUID of the group
	GroupId string `json:"group_id,omitempty"`
	// Array of group leaders and their attributes
	Leaders []ClusterGroupServiceLeader `json:"leaders,omitempty"`
	// Group status
	GroupStatus string `json:"group_status,omitempty"`
	// Type of the group
	GroupType string `json:"group_type,omitempty"`
}