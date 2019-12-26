/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type NatCounters struct {
	// The number of packets
	TotalPackets int64 `json:"total_packets,omitempty"`
	// The number of bytes
	TotalBytes int64 `json:"total_bytes,omitempty"`
	// The number of active sessions
	ActiveSessions int64 `json:"active_sessions,omitempty"`
}