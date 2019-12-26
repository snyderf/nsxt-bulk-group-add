/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LbCookieTime struct {
	// Both session cookie and persistence cookie are supported, Use LbSessionCookieTime for session cookie time setting, Use LbPersistenceCookieTime for persistence cookie time setting 
	Type_ string `json:"type"`
}
