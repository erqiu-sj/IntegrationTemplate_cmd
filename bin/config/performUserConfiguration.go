package config

import (
	"gopkg.in/ini.v1"
	"log"
)

func PerformUserConfiguration(result string, cfg *ini.File, allIniExecutableFileConfig map[string]RunIniFileConfigure) {
	if result == ADD_SECTION {
		return
	}
	log.Println(
		allIniExecutableFileConfig[result].Depository,
	)
	log.Println(
		allIniExecutableFileConfig[result].ExecuteBeforePulling,
	)
	log.Println(
		allIniExecutableFileConfig[result].AfterPulling,
	)

}
