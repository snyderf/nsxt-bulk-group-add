/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// license properties
type License struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// semicolon delimited feature list
	Features string `json:"features,omitempty"`
	// license edition
	Description string `json:"description,omitempty"`
	// product version
	ProductVersion string `json:"product_version,omitempty"`
	// date that license expires
	Expiry int64 `json:"expiry,omitempty"`
	// true for evalution license
	IsEval bool `json:"is_eval,omitempty"`
	// multi-hypervisor support
	IsMh bool `json:"is_mh,omitempty"`
	// license key
	LicenseKey string `json:"license_key,omitempty"`
	// whether the license has expired
	IsExpired bool `json:"is_expired,omitempty"`
	// product name
	ProductName string `json:"product_name,omitempty"`
	// License metrics specifying the capacity type of license key. Types are: - VM - CPU - USER(Concurrent User) 
	CapacityType string `json:"capacity_type,omitempty"`
	// license capacity; 0 for unlimited
	Quantity int64 `json:"quantity,omitempty"`
}
