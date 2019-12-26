/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type TraceflowObservationForwarded struct {
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
	// The name of the uplink the traceflow packet is forwarded on
	UplinkName string `json:"uplink_name,omitempty"`
	// The virtual tunnel endpoint label
	VtepLabel int64 `json:"vtep_label,omitempty"`
	// IP address of the destination end of the tunnel
	RemoteIpAddress string `json:"remote_ip_address,omitempty"`
	// The 64bit tunnel context carried on the wire
	Context int64 `json:"context,omitempty"`
	// IP address of the source end of the tunnel
	LocalIpAddress string `json:"local_ip_address,omitempty"`
	// This field will not be always available. Use remote_ip_address when this field is not set.
	DstTransportNodeId string `json:"dst_transport_node_id,omitempty"`
	// The name of the transport node to which the traceflow packet is forwarded
	DstTransportNodeName string `json:"dst_transport_node_name,omitempty"`
}