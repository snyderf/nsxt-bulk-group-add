/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MacAddressCsvListResult struct {
	// File name set by HTTP server if API  returns CSV result as a file.
	FileName string `json:"file_name,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Results []MacTableCsvRecord `json:"results,omitempty"`
}
