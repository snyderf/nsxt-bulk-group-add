/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// An NSService element that represents an ALG protocol
type AlgTypeNsService struct {
	// The specific type of NSServiceElement
	ResourceType string `json:"resource_type"`
	// The Application Layer Gateway (ALG) protocol. Please note, protocol NBNS_BROADCAST and NBDG_BROADCAST are  deprecated. Please use UDP protocol and create L4 Port Set type of service instead. 
	Alg string `json:"alg"`
	// The destination_port cannot be empty and must be a single value.
	DestinationPorts []string `json:"destination_ports"`
	// Source ports
	SourcePorts []string `json:"source_ports,omitempty"`
}
