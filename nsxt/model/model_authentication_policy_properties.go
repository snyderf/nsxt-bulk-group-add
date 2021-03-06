/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Configuration of authentication policies for the NSX node
type AuthenticationPolicyProperties struct {
	// Link to this resource
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// In order to trigger an account lockout, all authentication failures must occur in this time window. If the reset period expires, the failed login count is reset to zero. Only applies to NSX Manager nodes. Ignored on other node types.
	ApiFailedAuthResetPeriod int64 `json:"api_failed_auth_reset_period,omitempty"`
	// Minimum number of characters required in account passwords
	MinimumPasswordLength int64 `json:"minimum_password_length,omitempty"`
	// Once a lockout occurs, the account remains locked out of the CLI for this time period. While the lockout period is in effect, additional authentication attempts restart the lockout period, even if a valid password is specified.
	CliFailedAuthLockoutPeriod int64 `json:"cli_failed_auth_lockout_period,omitempty"`
	// Only applies to NSX Manager nodes. Ignored on other node types.
	ApiMaxAuthFailures int64 `json:"api_max_auth_failures,omitempty"`
	// Once a lockout occurs, the account remains locked out of the API for this time period. Only applies to NSX Manager nodes. Ignored on other node types.
	ApiFailedAuthLockoutPeriod int64 `json:"api_failed_auth_lockout_period,omitempty"`
	// Number of authentication failures that trigger CLI lockout
	CliMaxAuthFailures int64 `json:"cli_max_auth_failures,omitempty"`
}
