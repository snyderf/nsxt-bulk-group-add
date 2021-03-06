/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Directory domain synchronization statistics
type DirectoryDomainSyncStats struct {
	// Directory domain previous sync status. It could be one of the following two states.
	PrevSyncStatus string `json:"prev_sync_status,omitempty"`
	// All the historical full sync are counted in calculating the average full sync time in milliseconds.
	AvgFullSyncTime int64 `json:"avg_full_sync_time,omitempty"`
	// Directory domain previous sync type. It could be one of the following three states. Right after the directory domain is configured, this field is set to IDLE.
	PrevSyncType string `json:"prev_sync_type,omitempty"`
	// number of successful historical full sync initiated either by system or by API request.
	NumFullSync int64 `json:"num_full_sync,omitempty"`
	// Since what time the current state has begun. The time is expressed in millisecond epoch time.
	CurrentStateBeginTime int64 `json:"current_state_begin_time,omitempty"`
	// All the historical delta sync are counted in calculating the average delta sync time in milliseconds.
	AvgDeltaSyncTime int64 `json:"avg_delta_sync_time,omitempty"`
	// Directory domain previous sync status error if last status was failure.
	PrevSyncError string `json:"prev_sync_error,omitempty"`
	// Current running state of the directory domain in synchronization life cycle. It could be one of the following three states.
	CurrentState string `json:"current_state,omitempty"`
	// number of successful historical delta sync initiated either by system or by API request.
	NumDeltaSync int64 `json:"num_delta_sync,omitempty"`
	// Directory domain previous sync ending time expressed in millisecond epoch time.
	PrevSyncEndTime int64 `json:"prev_sync_end_time,omitempty"`
}
