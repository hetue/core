package kernel

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newContext, // 上下文
	).Build().Apply()
}
