/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Information about the upgrade bundle
type UpgradeBundleInfo struct {
	// URL for uploading upgrade bundle
	Url string `json:"url,omitempty"`
	// size of upgrade bundle
	BundleSize string `json:"bundle_size,omitempty"`
}
