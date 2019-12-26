/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type SecurityCertificate struct {
	// X.509 certificate in text form
	Text string `json:"text,omitempty"`
	// The time when the certificate starts being valid
	ValidFrom string `json:"valid_from,omitempty"`
	SshPublicKey string `json:"ssh_public_key,omitempty"`
	// The time when the certificate stops being valid
	ValidTo string `json:"valid_to,omitempty"`
	// The certificate must include the enclosing \"-----BEGIN CERTIFICATE-----\" and \"-----END CERTIFICATE-----\"
	PemEncoded string `json:"pem_encoded"`
}