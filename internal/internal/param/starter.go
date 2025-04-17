package param

import (
	"github.com/hetue/boot/internal/internal/internal/config"
)

type Bootstrap struct {
	Name      string
	Usage     string
	Copyright string
	Metadata  map[string]string
	Code      *config.Code
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{
		Name:      "hetu",
		Usage:     "河图持续集成系统插件，可使用`run`命令执行",
		Copyright: "成都睿景承科技有限公司",
		Metadata:  make(map[string]string),
		Code:      config.NewCode(),
	}
}
