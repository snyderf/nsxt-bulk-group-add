/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Alarm in NSX that needs users' attention
type Alarm struct {
	// Sources emitting this alarm
	Sources *AlarmSource `json:"sources,omitempty"`
	// Alarm state
	State string `json:"state"`
	// Alarm source component like nsx-manager, nsx-edge etc
	SourceComp string `json:"source_comp,omitempty"`
	// Severity of an Alarm
	Severity string `json:"severity,omitempty"`
	// Unique identifier(like UUID) for the node sending the Alarm
	SourceCompId string `json:"source_comp_id,omitempty"`
	// Date and time in UTC of the Alarm
	Timestamp int64 `json:"timestamp,omitempty"`
	// Description of the Alarm
	Message string `json:"message"`
	// Unique identifier for an NSX Alarm
	Id string `json:"id,omitempty"`
	// Alarm source sub component like nsx-mpa etc
	SourceSubcomp string `json:"source_subcomp,omitempty"`
}