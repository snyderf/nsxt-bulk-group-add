/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 2.5.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Guest virtual machine details include OS name and computer name of guest VM. 
type GuestInfo struct {
	// OS name of guest virtual machine. Currently this is supported for guests on ESXi that have VMware Tools installed. 
	OsName string `json:"os_name,omitempty"`
	// Computer name of guest virtual machine, which is set inside guest OS. Currently this is supported for guests on ESXi that have VMware Tools installed. 
	ComputerName string `json:"computer_name,omitempty"`
}