package command

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newRun,
	).Build().Apply()
}
