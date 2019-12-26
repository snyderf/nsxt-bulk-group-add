/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LbPoolMemberStatistics struct {
	// Pool member statistics counter
	Statistics *LbStatisticsCounter `json:"statistics"`
	// Pool member IP address
	IpAddress string `json:"ip_address"`
	// The port is configured in pool member. For virtual server port range case, pool member port must be null. 
	Port string `json:"port,omitempty"`
}