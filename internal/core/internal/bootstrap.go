package internal

import (
	"net/rpc"

	"github.com/hetue/core/internal/internal/internal/config"
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Lifecycle

	code *config.Code
	rpc  *rpc.Server
}

func NewBootstrap(bootstrap internal.Bootstrap) pangu.Bootstrap {
	return &Bootstrap{
		rpc:  bootstrap.Rpc,
		task: bootstrap.Task,
	}
}

func (b *Bootstrap) Startup(application *pangu.Application) error {
	return application.Add(b.rpc, b.task)
}
