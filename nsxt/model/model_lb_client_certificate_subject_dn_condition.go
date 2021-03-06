/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Match condition for client certficate subject DN
type LbClientCertificateSubjectDnCondition struct {
	// If true, case is significant when comparing subject DN value. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
	// Match type of subject DN
	MatchType string `json:"match_type,omitempty"`
	// Value of subject DN
	SubjectDn string `json:"subject_dn"`
}
