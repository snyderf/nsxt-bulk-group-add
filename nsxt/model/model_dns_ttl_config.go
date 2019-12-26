/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Firewall to use TTL config to manage firewall cache to expire snooped FQDN entries from DNS response.
type DnsTtlConfig struct {
	// The number of seconds that snooped DNS responses are retained in the cache. Used only when dns_ttl_type is USE_TTL.
	Ttl int64 `json:"ttl,omitempty"`
	// TTL type to decide how to manage ttl.
	DnsTtlType string `json:"dns_ttl_type"`
}
