package core

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newCommand, // 命令执行
	).Build().Apply()
}
