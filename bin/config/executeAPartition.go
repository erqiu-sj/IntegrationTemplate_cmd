package config

import (
	"IntegrationTemplate_cmd/bin/utils"
	"errors"
	"fmt"
)

// ExecuteAPartition 执行一个分区 准确的说是验证一个分区,并不是在这个函数中中执行
func ExecuteAPartition(conf *IniFileConfigure) {
	// 执行标记
	run := true
	run = runDepository(conf)
	runExecuteBeforePulling(conf, run)
	runAfterPulling(conf, run)
}

func runDepository(conf *IniFileConfigure) bool {
	if (!utils.VerifyPrefix(conf.Depository, "https://gitee.com") && !utils.VerifyPrefix(conf.Depository, "https://github.com")) && !utils.VerifyPrefix(conf.Depository, "default") {
		panic(errors.New(fmt.Sprint("depository must be a GitHub or Gitee address", " ===> ", conf.Depository)))
	}
	if utils.DefaultSkip(conf.Depository) {
		return true
	}
	//utils.Cmd("git", conf.Depository)
	return true
}
func runExecuteBeforePulling(conf *IniFileConfigure, runTag bool) {
	if runTag == false {
		return
	}
	utils.DivisionCmdRun(conf.ExecuteBeforePulling)
}
func runAfterPulling(conf *IniFileConfigure, runTag bool) {
	if runTag == false {
		return
	}
}
