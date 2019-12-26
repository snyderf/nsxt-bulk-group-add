package utilities

import (
	"nsxt-bulk-group-add/nsxt/model"
)

type NSXObjTags struct {
	Objtype     string       `yaml:"objtype"`
	DisplayName string       `yaml:"displayName"`
	Tags        []model.Tag  `yaml:"tags"`
}

type YamlNSXTags struct {
	NsxObject []NSXObjTags `yaml:"nsxObject"`
}


type VmApps struct {
	App 		string		`yaml:"app"`
	Tier	 	string 		`yaml:"tier"`
}

type VmTiers struct {
	Tier 		string		`yaml:"tier"`
	VM			string		`yaml:"vm"`
}

type IpApps struct{
	App 		string 		`yaml:"app"`
	IP 			string		`yaml:"ip"`
}