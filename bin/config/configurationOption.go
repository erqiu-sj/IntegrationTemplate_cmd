package config

import (
	"github.com/manifoldco/promptui"
	"gopkg.in/ini.v1"
)

type RunOptionsCallerFnGroup func(result string, filePtr *ini.File, allIniExecutableFileConfig map[string]RunIniFileConfigure)

// RunOptions 展示选项
func RunOptions(selectOption promptui.Select, fnGroup []RunOptionsCallerFnGroup, filePtr *ini.File, allIniExecutableFileConfig map[string]RunIniFileConfigure) {

	prompt := selectOption

	_, result, err := prompt.Run()

	if err != nil {
		panic(err)
		return
	}
	for _, fn := range fnGroup {
		fn(result, filePtr, allIniExecutableFileConfig)
	}
}
