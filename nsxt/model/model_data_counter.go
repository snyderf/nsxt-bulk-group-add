/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DataCounter struct {
	// The total packets or bytes
	Total int64 `json:"total"`
	// The multicast and broadcast packets or bytes
	MulticastBroadcast int64 `json:"multicast_broadcast,omitempty"`
	// The dropped packets or bytes
	Dropped int64 `json:"dropped,omitempty"`
}
