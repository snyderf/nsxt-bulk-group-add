/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// The settings are used during deployment and consequent update of an edge, unless indicated otherwise. The settings are preferred over the deprecated settings in VsphereDeploymentConfig. The settings reflect the current configuraton on an edge node. If the settings lag with actual state on the edge, these may be refreshed at NSX Manager using API POST api/v1/transport-nodes /&lt;transport-node-id&gt;?action=refresh_node_configuration&resource_type= EdgeNode 
type EdgeNodeSettings struct {
	// List of domain names that are used to complete unqualified host names. 
	SearchDomains []string `json:"search_domains,omitempty"`
	// List of DNS servers. 
	DnsServers []string `json:"dns_servers,omitempty"`
	// List of NTP servers. 
	NtpServers []string `json:"ntp_servers,omitempty"`
	// Host name or FQDN for edge node.
	Hostname string `json:"hostname,omitempty"`
	// Enabling SSH service is not recommended for security reasons. 
	EnableSsh bool `json:"enable_ssh,omitempty"`
	// Allowing root SSH logins is not recommended for security reasons. Edit of this property is not supported when updating transport node. Use the CLI to change this property. 
	AllowSshRootLogin bool `json:"allow_ssh_root_login,omitempty"`
}
