/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Preconfigured host switch specification is used for manually configured transport node. It is user's responsibility to ensure correct configuration is provided to NSX. This type is only valid for supported KVM fabric nodes.
type PreconfiguredHostSwitchSpec struct {
	ResourceType string `json:"resource_type"`
	// Preconfigured Transport Node host switches
	HostSwitches []PreconfiguredHostSwitch `json:"host_switches"`
}