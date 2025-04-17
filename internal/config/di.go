package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newWrapper,

		newRetry,   // 重试
		newRuntime, // 运行时
	).Build().Apply()
}
