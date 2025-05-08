package finder

import (
	"github.com/harluo/config"
	"github.com/harluo/di"
	"github.com/hetue/boot/internal/internal/internal/finder/internal"
)

func init() {
	di.New().Instance().Put(func(finder *internal.Default) config.Finder {
		return finder
	}).Group("finders").Build().Apply()
}
