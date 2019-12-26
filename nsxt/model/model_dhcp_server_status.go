/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DhcpServerStatus struct {
	// Error message, if available
	ErrorMessage string `json:"error_message,omitempty"`
	// UP means the dhcp service is working fine on both active transport-node and stand-by transport-node (if have), hence fail-over can work at this time if there is failure happens on one of the transport-node; DOWN means the dhcp service is down on both active transport-node and stand-by node (if have), hence the dhcp-service will not repsonse any dhcp request; Error means error happens on transport-node(s) or no status is reported from transport-node(s). The dhcp service may be working (or not working); NO_STANDBY means dhcp service is working in one of the transport node while not in the other transport-node (if have). Hence if the dhcp service in the working transport-node is down, fail-over will not happen and the dhcp service will go down. 
	ServiceStatus string `json:"service_status"`
	// uuid of stand_by transport node. null if non-HA mode
	StandByNode string `json:"stand_by_node,omitempty"`
	// uuid of active transport node
	ActiveNode string `json:"active_node"`
}