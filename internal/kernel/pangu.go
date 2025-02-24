package kernel

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newCommand,
	).Build().Apply()
}
