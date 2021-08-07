package main

import (
	"IntegrationTemplate_cmd/bin/config"
	"github.com/manifoldco/promptui"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	cfg, _ := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, "conf.ini")
	all := cfg.SectionStrings()
	for _, sectionName := range all {
		SectionConfig := new(config.IniFileConfigure)
		mappingStructureErr := cfg.Section(sectionName).MapTo(SectionConfig)
		if mappingStructureErr != nil {
			log.Fatalln("Cannot map to structure")
		}
		config.ExecuteAPartition(SectionConfig)
	}
	config.RunOptions(promptui.Select{
		Label: "选择你的操作",
		Items: []string{config.ADD_SECTION},
	}, []config.RunOptionsCallerFnGroup{
		config.AddSection,
	}, cfg)
}
