package get

import (
	"github.com/goexl/log"
	"github.com/hetue/core/internal/kernel"
	"github.com/pangum/pangu"
)

type Run struct {
	pangu.Get

	Steps  []kernel.Step
	Logger log.Logger
}
