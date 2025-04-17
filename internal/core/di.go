package core

import (
	"github.com/pangum/core"
)

func init() {
	core.New().Get().Dependency().Puts(
		newCommand, // 命令执行
	).Build().Apply()
}
