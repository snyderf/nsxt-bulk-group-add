/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ServiceInstanceStatus struct {
	// Deployment status of NXGI Partner Service-VM.
	InstanceDeploymentStatus *ServiceDeploymentStatus `json:"instance_deployment_status,omitempty"`
	// This object contains the list of issues which might come during post deployment configuration for a particular service instance.
	ConfigurationIssue *SvmConfigureIssue `json:"configuration_issue,omitempty"`
	// Id of an instantiation of a registered service.
	ServiceInstanceId string `json:"service_instance_id,omitempty"`
	// Health status of NXGI components on Partner Service-VM.
	InstanceHealthStatus *ServiceInstanceHealthStatus `json:"instance_health_status,omitempty"`
}
