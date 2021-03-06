/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Label that will be displayed for a UI element.
type Label struct {
	// Text to be displayed at the label.
	Text string `json:"text"`
	// If true, displays the label only on hover
	Hover bool `json:"hover,omitempty"`
	// Hyperlink of the specified UI page that provides details.
	Navigation string `json:"navigation,omitempty"`
	// If the condition is met then the label will be applied. Examples of expression syntax are provided under example_request section of CreateWidgetConfiguration API.
	Condition string `json:"condition,omitempty"`
	// Icons to be applied at dashboard for the label
	Icons []Icon `json:"icons,omitempty"`
}
