package main

import (
	"fmt"
	"io/ioutil"
	"nsxt-bulk-group-add/utilities"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// Need to include local packages as well as parameter input package.
/*
type importConfig struct {
	importConfig []nsxConfig `yaml:importconfig`
}
*/

type nsxConfig struct {
	Hostname   		string `yaml:"hostname"`
	Username   		string `yaml:"username"`
	Password   		string `yaml:"password"`
	VmImportfile 	string `yaml:"vmimportfile"`
	IpImportfile	string `yaml:"ipimportfile"`
}

func main() {
	//Need to validate that the config file is properly opened and parsed properly
	cfg := nsxConfig{}

	//Prepare regex parsers to determine the type of file that will be opened.
	//yamlType, _ := regexp.Compile(".*yaml$")
	csvType, _ := regexp.Compile(".*csv$")

	cfgsrc, err := ioutil.ReadFile("config.yaml")

	//check to see if there was an error in opening the yaml file
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", cfgsrc)
	//Convert from yaml file to go struct
	err = yaml.Unmarshal([]byte(cfgsrc), &cfg)

	fmt.Println(cfg)
	//validate that the config yaml was able to be converted properly or exit with error.
	if err != nil {
		panic(err)
	}

	//cfg := config.importConfig[0]

	//Validate that the filetype in the config file is a ".yaml" or ".csv" file and if so import the file to the nsxTags Variable.
	//fmt.Println("opening file: ", cfg.VMImportfile)

/* 
	**** Removing YAML functionality from first rev of the script

	if yamlType.MatchString(cfg.VmImportfile) {
		nsxTags = utilities.ImportYAMLTagsFromFile(cfg.Importfile)
	} else if csvType.MatchString(cfg.Importfile) {
		nsxTags = utilities.ImportCSVTagsFromFile(cfg.Importfile)
	} else {
		fmt.Println("The filename you have specified in the Config file does not mee the specific extensions supported by this program")
		fmt.Println("This program only supports \".yaml\" and \".csv\" file types at this time")
		fmt.Println("Please verify that your file is in the correct format and named properly.")
		panic("configuration error: file type in config.yaml:importfile invalid please correct.") //end the program if the file type is invalid
	}
*/

	//Validate the file format is a CSV File

		//Once the filetype has been validated and imported connect to the NSX manager
		nsxClient := utilities.ConnectNSXManager(cfg.Hostname, cfg.Username, cfg.Password)
		nsxPolicy := utilities.ConnectNSXPolicyManager(cfg.Hostname, cfg.Username, cfg.Password)

	if (csvType.MatchString(cfg.VmImportfile) && (cfg.VmImportfile != "none")) {
		
		vmAppConfig, vmTierConfig := utilities.GetVMAppsFromCSV(cfg.VmImportfile)
		
		tiers := utilities.AddVMTiers(nsxPolicy, nsxClient, vmTierConfig)
		utilities.AddAppGroups(nsxPolicy, vmAppConfig, tiers)

	} else if (cfg.VmImportfile == "none"){

		fmt.Println ("VM App config CSV not used skiping")

	} else {

		fmt.Println("VM Import  file is not a CSV file. Please verify and check again.")
		fmt.Println("If no file is to be added then please use \"none\" in the file to indicate is will not be used.")
		return
	}
	
	if (csvType.MatchString(cfg.IpImportfile) && (cfg.IpImportfile != "none")){

		ipAppConfig := utilities.GetIPAppsFromCSV(cfg.IpImportfile)

		utilities.AddIPGroups(nsxPolicy, ipAppConfig)
		fmt.Println("IP Groups have been added to the NSX Manager")

	}  else if (cfg.IpImportfile == "none"){

		fmt.Println ("IP App config CSV not used skiping")

	} else {
		fmt.Println("IPSet input file is not a CSV file. Please verify and check again.")
		fmt.Println("If no file is to be added then please use \"none\" in the file to indicate is will not be used.")
		return
	}

	
}
