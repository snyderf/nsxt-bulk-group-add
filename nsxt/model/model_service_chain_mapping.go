/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// A ServiceInsertionServiceProfile can be part of multiple ServiceChains. ServiceChainMapping for a particular profile will contain a list of all the ServiceChains it's part of. Each Mapping will also contain some metadata to uniquely identify a profile from other profiles.
type ServiceChainMapping struct {
	// A unique id generated for every ServiceChain. This is not a uuid.
	ServiceChainId string `json:"service_chain_id,omitempty"`
	// Each ServiceChain has forward_path_service_profiles and reverse_path_service_profiles. This property will indicate which of them being used. FORWARD - forward_path_service_profiles REVERSE - reverse_path_service_profiles
	Direction string `json:"direction,omitempty"`
	// Service Index represents a numerical position of a ServiceInsertionServiceProfile in a ServiceChain. It will be in reverse order. Service Index can point to either forward_path_service_profiles or reverse_path_service_profiles indicated by direction property. Example - For a ServiceChain A-B-C, A will have index of 3, B will have index of 2 and C will have index of 1.
	ServiceIndex int64 `json:"service_index,omitempty"`
}
