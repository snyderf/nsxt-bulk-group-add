/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Summarized view of all selected IPSec VPN sessions.
type IpSecVpnSessionSummary struct {
	// Traffic summary per session.
	TrafficSummaryPerSession []IpSecVpnSessionTrafficSummary `json:"traffic_summary_per_session,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Session summary for number of total, established, failed and degraded IPSec VPN sessions.
	SessionSummary *IPsecVpnikeSessionSummary `json:"session_summary,omitempty"`
	// Aggregate traffic statistics across all selected sessions.
	AggregateTrafficCounters *IpSecVpnTrafficCounters `json:"aggregate_traffic_counters,omitempty"`
}
