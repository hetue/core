package core

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newCommand, // 命令执行
	).Build().Apply()
}
