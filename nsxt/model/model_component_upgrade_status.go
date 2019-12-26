/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ComponentUpgradeStatus struct {
	// Upgrade status of component
	Status string `json:"status,omitempty"`
	// Pre-upgrade status of the component-type
	PreUpgradeStatus *UpgradeChecksExecutionStatus `json:"pre_upgrade_status,omitempty"`
	// Details about the upgrade status
	Details string `json:"details,omitempty"`
	// Component type for the upgrade status
	ComponentType string `json:"component_type,omitempty"`
	// Number of nodes of the type and at the component version
	NodeCountAtTargetVersion int32 `json:"node_count_at_target_version,omitempty"`
	// Target component version
	TargetComponentVersion string `json:"target_component_version,omitempty"`
	// Indicator of upgrade progress in percentage
	PercentComplete float32 `json:"percent_complete,omitempty"`
	// Can the upgrade of the remaining units in this component be skipped
	CanSkip bool `json:"can_skip,omitempty"`
	// Mapping of current versions of nodes and counts of nodes at the respective versions.
	CurrentVersionNodeSummary *NodeSummaryList `json:"current_version_node_summary,omitempty"`
}
