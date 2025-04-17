package get

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
	"github.com/hetue/boot/internal/config"
)

type Run struct {
	di.Get

	Retry  *config.Retry
	Logger log.Logger
}
