package utilities

import (
	"fmt"
	"nsxt-bulk-group-add/nsxt"
	"nsxt-bulk-group-add/nsxt/model"
)

func AddIPGroups(nsxPolicy *nsxt.APIClient, apps []IpApps) {
	
	var effectiveGroups []string
	var uniqueGroups[] string

	for _, app := range apps {
		effectiveGroups = append(effectiveGroups, app.App)
	}

	//Determine the unique Groups
	for _, eg := range effectiveGroups {
		skip := false
		for _, ug := range uniqueGroups {
			if eg == ug {
				skip = true
				break
			}
		}
		if !skip {
			uniqueGroups = append(uniqueGroups, eg)
		}
	}
	
	for _, ua := range uniqueGroups {
		var includedIPs []string

		for _, app := range apps {
			if (app.App == ua) {
				includedIPs = append(includedIPs,app.IP)	
			}
			
		}
				
		ipExp := []model.Expression{{
			IpAddresses: includedIPs,
			ResourceType: "IPAddressExpression",
		}}
	
		ipGroup := model.Group {
			DisplayName: ua,
			Expression: ipExp,
		}

		ipGroup, resp, err := nsxPolicy.PolicyInventoryGroupsApi.UpdateGroupForDomain(nsxPolicy.Context, "default", ua, ipGroup)
		if (err != nil) {
			fmt.Println(resp)
			panic(err)
		} else {
			fmt.Println("Group", ua, "Created Successfully")
		}
	}

	return
}

func AddVMTiers(nsxPolicy *nsxt.APIClient, nsxClient *nsxt.APIClient, vmTierConfig []VmTiers) []model.Group {
	var tiers []model.Group
	var effectiveTiers []string
	var uniqueTiers []string

	//Generate a list of the effective Tiers
	for _, vmTier := range vmTierConfig {
		effectiveTiers = append(effectiveTiers, vmTier.Tier)
	}
	
	//Determine the unique tiers 
	for _, et := range effectiveTiers {
		skip := false
		for _, ut := range uniqueTiers {
			if et == ut {
				skip = true
				break
			}
		}
		if !skip {
			uniqueTiers = append(uniqueTiers, et)
		}
	}
	//Generate a list of VMs and thier associated externalIDs
	vmList, resp, err := nsxClient.FabricVirtualMachinesApi.ListVirtualMachines(nsxClient.Context, nil)
	if (err != nil) {
		fmt.Println(resp)
		panic(err)
	}
	
	//Iterate through the unique Tier Instances and add VMs to the appropriate Tier
	for _, ut := range uniqueTiers {
		//Instantiate the extIDs variable during each iteration of the loop
		extIDs := []string{}
		for _, tierVM := range vmTierConfig {
			if tierVM.Tier == ut {
				for _, vm := range vmList.Results {
					fmt.Println(vm.DisplayName,tierVM.VM)
					if (tierVM.VM == vm.DisplayName) {
						fmt.Println(vm.DisplayName, vm.ExternalId)
						extIDs = append(extIDs, vm.ExternalId)
						break
					}
				}
			}
		}
		
		fmt.Println(extIDs)

		exp := []model.Expression{{
			ExternalIds: extIDs,
			MemberType: "VirtualMachine",
			ResourceType: "ExternalIDExpression",
		}}
	
	
		newGroup := model.Group {
			DisplayName: ut,
			Expression: exp,
		}
	
		newGroup, resp, err = nsxPolicy.PolicyInventoryGroupsApi.UpdateGroupForDomain(nsxPolicy.Context, "default", ut, newGroup)
		if (err != nil) {
			fmt.Println(resp)
			panic(err)
		}
		tiers = append(tiers, newGroup)
		fmt.Println("Tier", ut, "Added to NSXT")
	}

	return tiers
}

func AddAppGroups(nsxPolicy *nsxt.APIClient, vmAppConfig []VmApps, tiers []model.Group){
	var effectiveGroups []string
	var uniqueGroups []string

	//Generate a list of the effective Groups
	for _, vmApp := range vmAppConfig {
		effectiveGroups = append(effectiveGroups, vmApp.App)
	}

	//Determine the unique Groups
	for _, eg := range effectiveGroups {
		skip := false
		for _, ug := range uniqueGroups {
			if eg == ug {
				skip = true
				break
			}
		}
		if !skip {
			uniqueGroups = append(uniqueGroups, eg)
		}
	}

	//Assign the tiers to their parent app
	for _, ug := range uniqueGroups {
		var paths []string
		for _, appTier := range vmAppConfig {
			if appTier.App == ug {
				for _, ag := range tiers {
					if ag.Id == appTier.Tier {
						paths = append(paths, ag.Path)
					}
				}
			} 
		}
		exp := []model.Expression{{
			Paths: 			paths,
			ResourceType: 	"PathExpression",
		}}
	
	
		newGroup := model.Group {
			DisplayName: ug,
			Expression: exp,
		}

		newGroup, resp, err := nsxPolicy.PolicyInventoryGroupsApi.UpdateGroupForDomain(nsxPolicy.Context, "default", ug, newGroup)
		if (err != nil) {
			fmt.Println(resp)
			panic(err)
		}
		fmt.Println("Application", ug, "Added to NSXT")
	}
}
