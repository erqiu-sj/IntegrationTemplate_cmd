package config

// IniFileConfigure 每一段配置所需的选项 ! === 必填
type IniFileConfigure struct {
	Depository           string `ini:"depository"`           // !模版仓库地址
	ExecuteBeforePulling string `ini:"executeBeforePulling"` // 拉去前执行
	AfterPulling         string `ini:"afterPulling"`         // 仓库拉取后执行
}
