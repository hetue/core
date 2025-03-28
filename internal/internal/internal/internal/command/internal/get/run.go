package get

import (
	"github.com/goexl/log"
	"github.com/hetue/core/internal/config"
	"github.com/pangum/pangu"
)

type Run struct {
	pangu.Get

	Retry  *config.Retry
	Logger log.Logger
}
