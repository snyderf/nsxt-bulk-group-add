/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// This action is performed in HTTP forwarding phase. It is used to inspect the variable of HTTP request, and look up the persistence entry with its value and pool uuid as key. If the persistence entry is found, the HTTP request is forwarded to the recorded backend server according to the persistence entry. If the persistence entry is not found, a new entry is created in the table after backend server is selected. 
type LbVariablePersistenceOnAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// The property is used to enable a hash operation for variable value when composing the persistence key. 
	VariableHashEnabled bool `json:"variable_hash_enabled,omitempty"`
	// The property is the name of variable to be used. It specifies which variable's value of a HTTP Request will be used in the key of persistence entry. The variable can be a system embedded variable such as \"_cookie_JSESSIONID\", a customized variable defined in LbVariableAssignmentAction or a captured variable in regular expression such as \"article\". 
	VariableName string `json:"variable_name"`
	// If the persistence profile UUID is not specified, a default persistence table is created per virtual server. Currently, only LbGenericPersistenceProfile is supported. 
	PersistenceProfileId string `json:"persistence_profile_id,omitempty"`
}