/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Identity Firewall standalone hosts switch setting. This setting enables or disables Identity Firewall feature on all standalone hosts. 
type IdfwStandaloneHostsSwitchSetting struct {
	// IDFW standalone hosts switch (true=Enabled / false=Disabled).
	StandaloneHostsEnabled bool `json:"standalone_hosts_enabled"`
}
