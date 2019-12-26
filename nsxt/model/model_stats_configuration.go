/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Represents configuration of a statistic for an entity. Example, number of logical switches and their admin states.
type StatsConfiguration struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Title of the widget. If display_name is omitted, the widget will be shown without a title.
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// Supported visualization types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration and ContainerConfiguration.
	ResourceType string `json:"resource_type"`
	// The 'datasources' represent the sources from which data will be fetched. Currently, only NSX-API is supported as a 'default' datasource. An example of specifying 'default' datasource along with the urls to fetch data from is given at 'example_request' section of 'CreateWidgetConfiguration' API.
	Datasources []Datasource `json:"datasources,omitempty"`
	// Specify relavite weight in WidgetItem for placement in a view. Please see WidgetItem for details.
	Weight int32 `json:"weight,omitempty"`
	// Icons to be applied at dashboard for widgets and UI elements.
	Icons []Icon `json:"icons,omitempty"`
	// Please use the property 'shared' of View instead of this. The widgets of a shared view are visible to other users.
	Shared bool `json:"shared,omitempty"`
	Footer *Footer `json:"footer,omitempty"`
	// Id of drilldown widget, if any. Id should be a valid id of an existing widget. A widget is considered as drilldown widget when it is associated with any other widget and provides more detailed information about any data item from the parent widget.
	DrilldownId string `json:"drilldown_id,omitempty"`
	// Set to true if this widget should be used as a drilldown.
	IsDrilldown bool `json:"is_drilldown,omitempty"`
	// Legend to be displayed. If legend is not needed, do not include it.
	Legend *Legend `json:"legend,omitempty"`
	// Hyperlink of the specified UI page that provides details.
	Navigation string `json:"navigation,omitempty"`
	// Expression that fetches statistic. It can be used to show the characteristics of entities such as Logical Switches, Firewall Rules, and so on. For example, number of logical switches and their admin states. If stat is not provided, then it will not be displayed.
	Stat *StatItem `json:"stat,omitempty"`
	// Sections
	Sections []DonutSection `json:"sections,omitempty"`
	// A sub-type of StatsConfiguration. If sub-type is not specified the parent type is rendered. The COMPACT sub_type, conserves the space for the widget. The statistic is placed on the right side on top of the status bar and the title of the widget is placed on the left side on the top of the status bar. The COMPACT style aligns itself horizontally as per the width of the container. If multiple widgets are placed insided the container then the widgets are placed one below the other to conserve the space.
	SubType string `json:"sub_type,omitempty"`
	// Displayed at the sections, by default. It labels the entities of sections. If label is not provided, the sections are not labelled.
	Label *Label `json:"label,omitempty"`
}
