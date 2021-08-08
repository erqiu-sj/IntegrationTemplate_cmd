package config

import (
	"IntegrationTemplate_cmd/bin/utils"
	"errors"
	"fmt"
)

// ExecuteAPartition 执行一个分区 准确的说是验证一个分区,并不是在这个函数中中执行,不符标准则会报错
func ExecuteAPartition(conf *IniFileConfigure) RunIniFileConfigure {

	// 执行标记
	run := true
	run, depositoryURL := runDepository(conf)
	result := RunIniFileConfigure{
		Depository:           depositoryURL,
		ExecuteBeforePulling: runExecuteBeforePulling(conf, run),
		AfterPulling:         runAfterPulling(conf, run),
	}
	return result
}

func runDepository(conf *IniFileConfigure) (runTag bool, cmd string) {
	if (!utils.VerifyPrefix(conf.Depository, "https://gitee.com") && !utils.VerifyPrefix(conf.Depository, "https://github.com")) && !utils.VerifyPrefix(conf.Depository, "default") {
		panic(errors.New(fmt.Sprint("depository must be a GitHub or Gitee address", " ===> ", conf.Depository)))
	}
	if utils.DefaultSkip(conf.Depository) {
		return true, ""
	}
	return true, conf.Depository
}
func runExecuteBeforePulling(conf *IniFileConfigure, runTag bool) [][]string {
	var result [][]string
	if runTag == false {
		return result
	}
	result = utils.DivisionCmdRun(conf.ExecuteBeforePulling)
	return result
}
func runAfterPulling(conf *IniFileConfigure, runTag bool) [][]string {
	var result [][]string
	if runTag == false {
		return result
	}
	if utils.DefaultSkip(conf.Depository) {
		return result
	}
	result = utils.DivisionCmdRun(conf.AfterPulling)
	return result
}
