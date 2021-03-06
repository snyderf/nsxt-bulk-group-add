/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// An NSService element that represents an ethertype protocol
type EtherTypeNsService struct {
	// The specific type of NSServiceElement
	ResourceType string `json:"resource_type"`
	// Type of the encapsulated protocol
	EtherType int64 `json:"ether_type"`
}
