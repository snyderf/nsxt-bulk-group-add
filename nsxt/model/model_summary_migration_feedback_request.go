/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Minimal description of feedback requests.
type SummaryMigrationFeedbackRequest struct {
	// Indicates if a valid response already exist for this feedback request.
	Resolved bool `json:"resolved,omitempty"`
	// Name of this object in the source NSX endpoint.
	VObjectName string `json:"v_object_name,omitempty"`
	// Indicates if previous response was invalid. Please provide a valid response.
	Rejected bool `json:"rejected,omitempty"`
	// Identifier for this object in the source NSX endpoint.
	VObjectId string `json:"v_object_id,omitempty"`
	// If the feedback request was resolved earlier, provides details about the previous resolution.
	Resolution string `json:"resolution,omitempty"`
	// Identifier of the feedback request.
	Id string `json:"id,omitempty"`
	// Identifier of the object for which feedback is requested.
	ObjectId string `json:"object_id,omitempty"`
	// Details about this specific feedback request.
	Details string `json:"details,omitempty"`
}