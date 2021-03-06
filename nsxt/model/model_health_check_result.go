/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Result of health check .
type HealthCheckResult struct {
	// Status of VLAN-MTU health check result; TRUNKED - all specified VLAN IDs are allowed by VLAN and MTU settings; UNTRUNKED - some/all specified VLAN IDs may be disallowed by VLAN or MTU settings; UNKNOWN - some/all health check result are unknown due to infrastructure issues. 
	VlanMtuStatus string `json:"vlan_mtu_status,omitempty"`
	// List of health check results on specific transport node 
	ResultsPerTransportNode []HealthCheckResultPerTransportNode `json:"results_per_transport_node,omitempty"`
}
