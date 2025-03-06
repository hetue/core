package config

import (
	"github.com/goexl/gox"
)

type Runtime struct {
	// 是否启用默认配置
	Default *bool `default:"true"`
	// 是否显示详细信息
	Verbose bool `default:"false"`
	// 是否在出错时打印输出
	Pwe *bool `default:"true"`

	_ gox.Pointerized
}

func newRuntime(wrapper *Wrapper) *Runtime {
	return wrapper.Runtime
}
