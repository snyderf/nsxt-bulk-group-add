/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Configuration for automatically joining a cluster node to the cluster after it is deployed. ClusteringConfig is required if any of the deployment nodes has CONTROLLER role. 
type ClusteringConfig struct {
	// Specifies the type of clustering config to be used. 
	ClusteringType string `json:"clustering_type"`
}
