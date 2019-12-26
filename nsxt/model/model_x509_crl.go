/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// A CRL is a time-stamped list identifying revoked certificates.
type X509Crl struct {
	// Next update time for the CRL
	NextUpdate string `json:"next_update,omitempty"`
	// CRL's version number either 1 or 2
	Version string `json:"version,omitempty"`
	// list of X509CrlEntry
	CrlEntries []X509CrlEntry `json:"crl_entries,omitempty"`
	// Issuer's distinguished name(DN)
	Issuer string `json:"issuer,omitempty"`
}
