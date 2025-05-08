package core

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newCommand, // 命令执行
	).Build().Apply()
}
