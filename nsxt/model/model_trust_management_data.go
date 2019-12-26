/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type TrustManagementData struct {
	// list of supported algorithms
	SupportedAlgorithms []CryptoAlgorithm `json:"supported_algorithms,omitempty"`
}