/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Upload status of upgrade bundle uploaded from url
type UpgradeBundleUploadStatus struct {
	// URL for uploading upgrade bundle
	Url string `json:"url,omitempty"`
	// Current status of upgrade bundle upload
	Status string `json:"status,omitempty"`
	// Detailed status of upgrade bundle upload
	DetailedStatus string `json:"detailed_status,omitempty"`
	// Percent of bundle uploaded from URL
	Percent float32 `json:"percent,omitempty"`
}