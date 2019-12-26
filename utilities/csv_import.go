package utilities

import (
	"encoding/csv"
	"fmt"
	"os"

	"nsxt-bulk-group-add/nsxt/model"
)

func ImportCSVTagsFromFile(filename string) []NSXObjTags {

	var config []NSXObjTags

	src, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	lines, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", lines)

	for _, line := range lines {

		lineTags := []model.Tag{{Scope: line[2], Tag: line[3]}}

		lineconfig := NSXObjTags{
			Objtype:     line[0],
			DisplayName: line[1],
			Tags:        lineTags,
		}

		config = append(config, lineconfig)
	}

	//fmt.Println("value: ", config.NsxObject[0].Tags)

	return config
}

func GetVMAppsFromCSV(filename string) ([]VmApps, []VmTiers) {

	var appConfig []VmApps
	var tierConfig []VmTiers

	src, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	lines, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", lines)

	for _, line := range lines {

		lineAppConfig := VmApps{
			App: 	line[0],
			Tier: 	line[1],
		}

		lineTierConfig := VmTiers{
			Tier: 	line[1],
			VM:		line[2],
		}
		appConfig = append(appConfig, lineAppConfig)
		tierConfig = append(tierConfig, lineTierConfig)
	}

	//fmt.Println("value: ", config.NsxObject[0].Tags)

	return appConfig, tierConfig
}

func GetIPAppsFromCSV(filename string) ([]IpApps) {

	var appConfig []IpApps

	src, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	lines, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", lines)

	for _, line := range lines {

		lineAppConfig := IpApps{
			App: 	line[0],
			IP: 	line[1],
		}
		
		appConfig = append(appConfig, lineAppConfig)
	}

	//fmt.Println("value: ", config.NsxObject[0].Tags)

	return appConfig
}