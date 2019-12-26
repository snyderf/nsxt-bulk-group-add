/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Answer of nslookup
type DnsQueryAnswer struct {
	// Unparsed answer string from raw_answer. 
	RawString string `json:"raw_string,omitempty"`
	// Matched name of the given address. 
	Name string `json:"name,omitempty"`
	// Can be resolved ip address. 
	Address string `json:"address,omitempty"`
}