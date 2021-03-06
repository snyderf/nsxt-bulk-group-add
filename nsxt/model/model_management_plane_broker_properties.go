/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Information about a management plane node this controller is configured to communicate with
type ManagementPlaneBrokerProperties struct {
	// IP address or hostname of the message bus broker on the management plane node.
	Host string `json:"host"`
	// Port number of the message bus broker on the management plane node.
	Port int64 `json:"port,omitempty"`
	// Certificate thumbprint of the message bus broker on the management plane node.
	Thumbprint string `json:"thumbprint"`
}
