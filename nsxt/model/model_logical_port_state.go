/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Contians realized state of the logical port. For example, transport node on which the port is located, discovered and realized address bindings of the logical port. 
type LogicalPortState struct {
	// If any address binding discovered on the port is also found on other port on the same logical switch, then it is included in the duplicate bindings list along with the ID of the port with which it conflicts. 
	DuplicateBindings []DuplicateAddressBindingEntry `json:"duplicate_bindings,omitempty"`
	// Identifiers of the transport node where the port is located
	TransportNodeIds []string `json:"transport_node_ids,omitempty"`
	// Contains the list of address bindings for a logical port that were automatically dicovered using various snooping methods like ARP, DHCP etc. 
	DiscoveredBindings []AddressBindingEntry `json:"discovered_bindings,omitempty"`
	// Id of the logical port
	Id string `json:"id"`
	// List of logical port bindings that are realized. This list may be populated from the discovered bindings or manual user specified bindings. This binding configuration can be used by features such as firewall, spoof-guard, traceflow etc. 
	RealizedBindings []AddressBindingEntry `json:"realized_bindings,omitempty"`
}
