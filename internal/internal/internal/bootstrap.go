package internal

import (
	"github.com/hetue/core/internal/internal/internal/internal"
	"github.com/hetue/core/internal/internal/internal/internal/command"
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Lifecycle

	run *command.Run
}

func NewBootstrap(bootstrap internal.Bootstrap) pangu.Bootstrap {
	return &Bootstrap{
		run: bootstrap.Run,
	}
}

func (b *Bootstrap) Startup(application *pangu.Application) error {
	return application.Add(b.run)
}
