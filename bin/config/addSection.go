package config

import (
	"errors"
	"github.com/manifoldco/promptui"
	"gopkg.in/ini.v1"
	"log"
)

func ValidateSectionName(result string) error {
	if result == "" {
		return errors.New("新的分区名不能为空")
	}
	return nil
}
func AddSection(result string, filePtr *ini.File, allIniExecutableFileConfig map[string]RunIniFileConfigure) {
	if result != ADD_SECTION {
		return
	}

	prompt := promptui.Prompt{Label: "给新的分区tag取个名吧?", Validate: ValidateSectionName}
	promptResult, runErr := prompt.Run()
	if runErr != nil {
		log.Println(runErr)
	}
	_, createErr := filePtr.NewRawSection(promptResult, "depository           = default\nexecuteBeforePulling = default\nafterPulling         = default\n")
	if createErr != nil {
		log.Println("创建失败", createErr)
	}
	log.Println("创建成功")
	saveErr := filePtr.SaveTo("conf.ini")
	if saveErr != nil {
		log.Println("保存配置失败，请重试")
	}
}
