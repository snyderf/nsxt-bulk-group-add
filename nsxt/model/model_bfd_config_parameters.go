/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// BFD configuration for the given Peer.
type BfdConfigParameters struct {
	// The time interval (in milliseconds) between heartbeat packets for BFD when receiving heartbeats.| For edge cluster type of bare metal, this value should be >= 300ms.| For edge cluster type of virtual machine or hybrid, this value should be >= 1000ms.
	ReceiveInterval int64 `json:"receive_interval,omitempty"`
	// Number of times a packet is missed before BFD declares the neighbor down.
	DeclareDeadMultiple int64 `json:"declare_dead_multiple,omitempty"`
	// The time interval (in milliseconds) between heartbeat packets for BFD when sending heartbeats.| For edge cluster type of bare metal, this value should be >= 300ms.| For edge cluster type of virtual machine or hybrid, this value should be >= 1000ms.
	TransmitInterval int64 `json:"transmit_interval,omitempty"`
}
