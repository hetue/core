package kernel

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newContext, // 上下文
	).Build().Apply()
}
