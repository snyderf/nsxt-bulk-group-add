/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AggregatedDataCounter struct {
	TxBytes *DataCounter `json:"tx_bytes,omitempty"`
	RxPackets *DataCounter `json:"rx_packets,omitempty"`
	TxPackets *DataCounter `json:"tx_packets,omitempty"`
	RxBytes *DataCounter `json:"rx_bytes,omitempty"`
}
