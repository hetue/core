package kernel

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newContext, // 上下文
	).Build().Apply()
}
