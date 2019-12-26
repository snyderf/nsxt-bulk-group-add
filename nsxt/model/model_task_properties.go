/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Task properties
type TaskProperties struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Current status of the task
	Status string `json:"status,omitempty"`
	// True if response for asynchronous request is available
	AsyncResponseAvailable bool `json:"async_response_available,omitempty"`
	// Description of the task
	Description string `json:"description,omitempty"`
	// The start time of the task in epoch milliseconds
	StartTime int64 `json:"start_time,omitempty"`
	// True if this task can be canceled
	Cancelable bool `json:"cancelable,omitempty"`
	// HTTP request method
	RequestMethod string `json:"request_method,omitempty"`
	// Name of the user who created this task
	User string `json:"user,omitempty"`
	// Task progress if known, from 0 to 100
	Progress int64 `json:"progress,omitempty"`
	// A message describing the disposition of the task
	Message string `json:"message,omitempty"`
	// URI of the method invocation that spawned this task
	RequestUri string `json:"request_uri,omitempty"`
	// Identifier for this task
	Id string `json:"id,omitempty"`
	// The end time of the task in epoch milliseconds
	EndTime int64 `json:"end_time,omitempty"`
}