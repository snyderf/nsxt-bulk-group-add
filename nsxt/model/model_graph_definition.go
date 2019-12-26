/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Defines a graph
type GraphDefinition struct {
	// Defines the points of a graph.
	PointDefinition *PointDefinition `json:"point_definition"`
	// Describes the graph. It labels the entities of graph. If the label is not provided then it is not shown for a graph. For example, for a single graph, the title of widget can describe the graph and a label may not be necessary to be shown.
	Label *Label `json:"label,omitempty"`
	// Additional rendering or conditional evaluation of the field values to be performed, if any.
	RenderConfiguration []RenderConfiguration `json:"render_configuration,omitempty"`
}
