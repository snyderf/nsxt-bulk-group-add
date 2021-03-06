/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Properties of the layout of a container or widget
type LayoutProperties struct {
	// Describes the number of rows of grid layout of a container or widget. This property is applicable for grid layout only.
	NumRows int32 `json:"num_rows,omitempty"`
	// Describes the number of columns of grid layout of a container or widget. This property is applicable for grid layout only.
	NumColumns int32 `json:"num_columns,omitempty"`
}
