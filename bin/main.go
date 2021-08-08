package main

import (
	"IntegrationTemplate_cmd/bin/config"
	"github.com/manifoldco/promptui"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	// 初始化可执行的map配置
	var allIniExecutableFileConfig map[string]config.RunIniFileConfigure
	allIniExecutableFileConfig = make(map[string]config.RunIniFileConfigure)
	// 初始化所有option选项
	optionItem := []string{config.ADD_SECTION}
	// 读取ini文件
	cfg, _ := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, "conf.ini")
	// 获取所有sectionName数组
	all := cfg.SectionStrings()
	// 遍历
	for _, sectionName := range all {
		SectionConfig := new(config.IniFileConfigure)
		mappingStructureErr := cfg.Section(sectionName).MapTo(SectionConfig)
		if mappingStructureErr != nil {
			log.Fatalln("Cannot map to structure")
		}
		newlyFiredConfiguration :=
			config.ExecuteAPartition(SectionConfig)
		if newlyFiredConfiguration.Depository != "" {
			optionItem = append(optionItem, sectionName)
			allIniExecutableFileConfig[sectionName] = newlyFiredConfiguration
		}
	}
	config.RunOptions(promptui.Select{
		Label: "选择你的操作",
		Items: optionItem,
	}, []config.RunOptionsCallerFnGroup{
		config.AddSection,
		config.PerformUserConfiguration,
	}, cfg, allIniExecutableFileConfig)
}
