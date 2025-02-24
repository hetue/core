package kernel

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newContext, // 上下文
		newCommand, // 命令执行
	).Build().Apply()
}
