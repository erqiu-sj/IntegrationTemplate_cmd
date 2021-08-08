package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func Cmd(name string, arg string) {
	cmd := exec.Command(name, arg)
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer func(stdout io.ReadCloser) {
		_ = stdout.Close()
	}(stdout)
	runErr := cmd.Start()
	if runErr != nil {
		log.Println(runErr)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}

// DivisionCmdRun 分割终端执行命令行字符串，转为一个二维数组
func DivisionCmdRun(cmd string) [][]string {
	var result [][]string
	if DefaultSkip(cmd) {
		return result
	}
	trimSpace := strings.TrimSpace(cmd)
	allCommands := strings.Split(trimSpace, "& ")
	for _, command := range allCommands {
		commandItem := strings.SplitN(command, " ", 2)
		result = append(result, commandItem)
	}
	return result
}
