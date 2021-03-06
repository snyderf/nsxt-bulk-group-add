/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Detailed feedback requests from the migration tool where user input is required.
type MigrationFeedbackRequest struct {
	// Indicates if a valid response already exist for this feedback request.
	Resolved bool `json:"resolved,omitempty"`
	// List of acceptable actions for this feedback request.
	AcceptedActions []string `json:"accepted_actions,omitempty"`
	// Identify a feedback request type across objects. This can be used to group together objects with similar feedback request and resolve them in one go.
	Hash string `json:"hash,omitempty"`
	// Functional area that this query falls into.
	Vertical string `json:"vertical,omitempty"`
	// Identifier for this object in the source NSX endpoint.
	VObjectId string `json:"v_object_id,omitempty"`
	// The suggested value to resolve this feedback request.
	SuggestedValue string `json:"suggested_value,omitempty"`
	// Detailed feedback request with options.
	Message string `json:"message,omitempty"`
	// Identifier of the object for which feedback is requested.
	ObjectId string `json:"object_id,omitempty"`
	// Data type of the items listed in acceptable values list.
	AcceptedValueType string `json:"accepted_value_type,omitempty"`
	// Name of this object in the source NSX endpoint.
	VObjectName string `json:"v_object_name,omitempty"`
	// Indicates if multiple values can be selected as response from the list of acceptable value.
	MultiValue bool `json:"multi_value,omitempty"`
	// List of acceptable values for this feedback request.
	AcceptedValues []string `json:"accepted_values,omitempty"`
	// Identifier of the feedback request.
	Id string `json:"id,omitempty"`
	// The suggested action to resolve this feedback request.
	SuggestedAction string `json:"suggested_action,omitempty"`
	// Functional sub-area that this query falls into.
	SubVertical string `json:"sub_vertical,omitempty"`
	// If the feedback request was resolved earlier, provides details about the previous resolution.
	Resolution string `json:"resolution,omitempty"`
	// Indicates if previous response was invalid. Please provide a valid response.
	Rejected bool `json:"rejected,omitempty"`
}
