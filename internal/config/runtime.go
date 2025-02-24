package config

type Runtime struct {
	// 是否启用默认配置
	Default *bool `default:"true"`
	// 是否在出错时打印输出
	Pwe *bool `default:"true"`
}

func newRuntime(wrapper *Wrapper) *Runtime {
	return wrapper.Runtime
}
