/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Duplicate address binding information
type DuplicateAddressBindingEntry struct {
	// Source from which the address binding entry was obtained
	Source string `json:"source,omitempty"`
	// Combination of IP-MAC-VLAN binding
	Binding *PacketAddressClassifier `json:"binding,omitempty"`
	// Timestamp at which the binding was discovered via snooping or manually specified by the user 
	BindingTimestamp int64 `json:"binding_timestamp,omitempty"`
	// Provides the ID of the port on which the same address bidning exists 
	ConflictingPort string `json:"conflicting_port,omitempty"`
}
