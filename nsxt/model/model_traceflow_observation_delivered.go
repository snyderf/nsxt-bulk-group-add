/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type TraceflowObservationDelivered struct {
	// Timestamp when the observation was created by the transport node (microseconds epoch)
	TimestampMicro int64 `json:"timestamp_micro,omitempty"`
	// The sub type of the component that issued the observation.
	ComponentSubType string `json:"component_sub_type,omitempty"`
	ResourceType string `json:"resource_type"`
	// The type of the component that issued the observation.
	ComponentType string `json:"component_type,omitempty"`
	// name of the transport node that observed a traceflow packet
	TransportNodeName string `json:"transport_node_name,omitempty"`
	// Timestamp when the observation was created by the transport node (milliseconds epoch)
	Timestamp int64 `json:"timestamp,omitempty"`
	// id of the transport node that observed a traceflow packet
	TransportNodeId string `json:"transport_node_id,omitempty"`
	// the hop count for observations on the transport node that a traceflow packet is injected in will be 0. The hop count is incremented each time a subsequent transport node receives the traceflow packet. The sequence number of 999 indicates that the hop count could not be determined for the containing observation.
	SequenceNo int64 `json:"sequence_no,omitempty"`
	// type of the transport node that observed a traceflow packet
	TransportNodeType string `json:"transport_node_type,omitempty"`
	// The name of the component that issued the observation.
	ComponentName string `json:"component_name,omitempty"`
	// This field specifies the resolution type of ARP ARP_SUPPRESSION_PORT_CACHE - ARP request is suppressed by port DB ARP_SUPPRESSION_TABLE - ARP request is suppressed by ARP table ARP_SUPPRESSION_CP_QUERY - ARP request is suppressed by info derived from CP ARP_VM - No suppression and the ARP request is resolved.
	ResolutionType string `json:"resolution_type,omitempty"`
	// The name of the logical port into which the traceflow packet was delivered
	LportName string `json:"lport_name,omitempty"`
	// The source MAC address of form: \"^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$\". For example: 00:00:00:00:00:00. 
	TargetMac string `json:"target_mac,omitempty"`
	// VLAN on bridged network
	VlanId int64 `json:"vlan_id,omitempty"`
	// The id of the logical port into which the traceflow packet was delivered
	LportId string `json:"lport_id,omitempty"`
}
