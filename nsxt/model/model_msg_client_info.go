/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Information about a messaging client
type MsgClientInfo struct {
	// Software version of the node.
	SoftwareVersion string `json:"software_version,omitempty"`
	// Account name for the messaging client. Identifies the client on the management plane message bus.
	AccountName string `json:"account_name,omitempty"`
	// Messaging client's certificate. Used to authenticate to the control plane messaging endpoint.
	Certificate string `json:"certificate,omitempty"`
	// Messaging client's secret. Used to authenticate to the management plane messaging broker.
	SharedSecret string `json:"shared_secret,omitempty"`
}
