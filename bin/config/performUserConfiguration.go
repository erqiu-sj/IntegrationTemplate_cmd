package config

import (
	"IntegrationTemplate_cmd/bin/utils"
	"gopkg.in/ini.v1"
)

func PerformUserConfiguration(result string, cfg *ini.File, allIniExecutableFileConfig map[string]RunIniFileConfigure) {

	if result == ADD_SECTION {
		return
	}

	utils.RunDepository(allIniExecutableFileConfig[result].Depository)
	utils.RunExecuteBeforePulling(allIniExecutableFileConfig[result].ExecuteBeforePulling)
	utils.RunAfterPulling(allIniExecutableFileConfig[result].AfterPulling)
}
