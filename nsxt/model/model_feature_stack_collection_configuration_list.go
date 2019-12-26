/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Feature stack collection configuration list result
type FeatureStackCollectionConfigurationList struct {
	// The complete set of feature stack data collection configurations
	Results []FeatureStackCollectionConfiguration `json:"results,omitempty"`
}
