package config

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newWrapper,

		newLogging, // 日志
		newProxy,   // 代理
		newRetry,   // 重试
		newRuntime, // 运行时
	).Build().Apply()
}
