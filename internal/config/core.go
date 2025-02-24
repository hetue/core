package config

type Logging struct {
	// 是否显示详细信息
	Verbose bool `default:"false"`
	// 日志配置
	Level string `default:"info"`
	// 是否在出错时打印输出
	Pwe *bool `default:"true"`
}
