/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type TraceflowObservationReplicationLogical struct {
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
	// The label of VTEP
	VtepLabel int64 `json:"vtep_label,omitempty"`
	// This field specifies the type of replication message TX_VTEP - Transmit replication to all VTEPs TX_MTEP - Transmit replication to all MTEPs RX - Receive replication
	ReplicationType string `json:"replication_type,omitempty"`
	// Local IP address of the component that replicates the packet.
	LocalIpAddress string `json:"local_ip_address,omitempty"`
	// The name of uplink
	UplinkName string `json:"uplink_name,omitempty"`
}